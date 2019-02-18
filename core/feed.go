package core

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/textileio/textile-go/pb"
)

var flatFeedTypes = []pb.Block_BlockType{
	pb.Block_JOIN,
	pb.Block_LEAVE,
	pb.Block_FILES,
	pb.Block_MESSAGE,
	pb.Block_COMMENT,
	pb.Block_LIKE,
}

var annotatedFeedTypes = []pb.Block_BlockType{
	pb.Block_JOIN,
	pb.Block_LEAVE,
	pb.Block_FILES,
	pb.Block_MESSAGE,
}

type feedStack struct {
	id       string
	top      *pb.Block
	children []*pb.Block
}

type feedItemOpts struct {
	annotations bool
	comments    []*pb.Comment
	likes       []*pb.Like
	target      *pb.FeedItem
}

func (t *Textile) Feed(offset string, limit int, threadId string, mode pb.FeedMode) (*pb.FeedItemList, error) {
	var types []pb.Block_BlockType
	switch mode {
	case pb.FeedMode_CHRONO, pb.FeedMode_STACKS:
		types = flatFeedTypes
	case pb.FeedMode_ANNOTATED:
		types = annotatedFeedTypes
	}

	var query string
	for i, t := range types {
		query += fmt.Sprintf("type=%d", t)
		if i != len(types)-1 {
			query += " or "
		}
	}
	query = "(" + query + ")"
	if threadId != "" {
		if t.Thread(threadId) == nil {
			return nil, ErrThreadNotFound
		}
		query = fmt.Sprintf("(threadId='%s') and %s", threadId, query)
	}

	blocks := t.Blocks(offset, limit, query)
	list := make([]*pb.FeedItem, 0)
	var count int

	switch mode {
	case pb.FeedMode_CHRONO, pb.FeedMode_ANNOTATED:
		for _, block := range blocks.Items {
			item, err := t.feedItem(block, feedItemOpts{
				annotations: mode == pb.FeedMode_ANNOTATED,
			})
			if err != nil {
				return nil, err
			}
			list = append(list, item)
			count++
		}

	case pb.FeedMode_STACKS:
		stacks := make([]feedStack, 0)
		var last *feedStack
		for _, block := range blocks.Items {
			if len(stacks) > 0 {
				last = &stacks[len(stacks)-1]
			}
			targetId := getTargetId(block)

			if len(stacks) == 0 || targetId != getTargetId(last.top) {
				// start a new stack
				stacks = append(stacks, feedStack{id: targetId, top: block})
			} else {
				// append to last
				last.children = append(last.children, block)
			}
		}

		for _, stack := range stacks {
			item, err := t.feedStackItem(stack)
			if err != nil {
				return nil, err
			}
			if item == nil {
				continue
			}
			list = append(list, item)
			count += len(stack.children) + 1
		}
	}

	var nextOffset string
	if len(blocks.Items) > 0 {
		nextOffset = blocks.Items[len(blocks.Items)-1].Id

		// see if there's actually more
		if len(t.datastore.Blocks().List(nextOffset, 1, query).Items) == 0 {
			nextOffset = ""
		}
	}

	return &pb.FeedItemList{
		Items: list,
		Count: int32(count),
		Next:  nextOffset,
	}, nil
}

func (t *Textile) feedItem(block *pb.Block, opts feedItemOpts) (*pb.FeedItem, error) {
	item := &pb.FeedItem{
		Block:   block.Id,
		Thread:  block.Thread,
		Payload: &any.Any{},
	}

	var payload proto.Message
	var err error
	switch block.Type {
	case pb.Block_JOIN:
		item.Payload.TypeUrl = "/Join"
		payload, err = t.join(block, opts)
	case pb.Block_LEAVE:
		item.Payload.TypeUrl = "/Leave"
		payload, err = t.leave(block, opts)
	case pb.Block_FILES:
		item.Payload.TypeUrl = "/Files"
		payload, err = t.file(block, opts)
	case pb.Block_MESSAGE:
		item.Payload.TypeUrl = "/Text"
		payload, err = t.message(block, opts)
	case pb.Block_COMMENT:
		item.Payload.TypeUrl = "/Comment"
		payload, err = t.comment(block, opts)
	case pb.Block_LIKE:
		item.Payload.TypeUrl = "/Like"
		payload, err = t.like(block, opts)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	value, err := proto.Marshal(payload)
	if err != nil {
		return nil, err
	}
	item.Payload.Value = value

	return item, nil
}

func (t *Textile) feedStackItem(stack feedStack) (*pb.FeedItem, error) {
	var comments []*pb.Comment
	var likes []*pb.Like

	// Does the stack contain the initial target,
	// or is it a continuation stack of just annotations?
	// We'll need to load the target in the latter case.
	var target *pb.Block
	handleChild := func(child *pb.Block) error {
		switch child.Type {
		case pb.Block_COMMENT:
			comment, err := t.comment(child, feedItemOpts{annotations: true})
			if err != nil {
				return err
			}
			comments = append(comments, comment)
		case pb.Block_LIKE:
			like, err := t.like(child, feedItemOpts{annotations: true})
			if err != nil {
				return err
			}
			likes = append(likes, like)
		default:
			target = child
		}
		return nil
	}
	for _, child := range stack.children {
		if err := handleChild(child); err != nil {
			return nil, err
		}
	}

	var initial bool
	if target != nil { // target was in children, newer annotations may exist, make target top
		initial = true
		if err := handleChild(stack.top); err != nil {
			return nil, err
		}
	} else if !isAnnotation(stack.top) { // target is top, newer annotations may exist
		initial = true
		target = stack.top
	} else { // target needs to be loaded, older annotations may exist
		if t.blockIgnored(stack.id) {
			return nil, nil
		}
		target = t.datastore.Blocks().Get(stack.id)
		if target == nil {
			return nil, nil
		}
	}

	targetItem, err := t.feedItem(target, feedItemOpts{
		comments: comments,
		likes:    likes,
	})
	if err != nil {
		return nil, err
	}

	if !initial {
		// target gets wrapped with the top block
		return t.feedItem(stack.top, feedItemOpts{
			target: targetItem,
		})
	}

	return targetItem, nil
}

func (t *Textile) blockIgnored(blockId string) bool {
	query := "target='ignore-" + blockId + "'"
	return len(t.datastore.Blocks().List("", -1, query).Items) > 0
}

func getTargetId(block *pb.Block) string {
	switch block.Type {
	case pb.Block_COMMENT, pb.Block_LIKE:
		return block.Target
	default:
		return block.Id
	}
}

func isAnnotation(block *pb.Block) bool {
	switch block.Type {
	case pb.Block_COMMENT, pb.Block_LIKE:
		return true
	default:
		return false
	}
}
