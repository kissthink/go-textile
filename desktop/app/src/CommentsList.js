import React, { Component } from 'react'
import { Divider, Header, Comment, Image } from 'semantic-ui-react'
import Moment from 'react-moment'
import MessageForm from './MessageForm'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class CommentsList extends Component {
  scrollToBottom () {
    this.target.scrollIntoView(false)
  }
  render () {
    const { store } = this.props
    return (
      <Comment.Group size='small'>
        {store.thread.currentUpdate && this.renderItem(store.thread.currentUpdate)}
        <Divider horizontal>
          <Header as='h6'>Replies</Header>
        </Divider>
        {store.thread.currentUpdate && store.thread.currentUpdate.comments
          .slice()
          .reverse()
          .map(comment => this.renderItem(comment))}
        <div style={{ height: '1em' }} ref={el => { this.target = el }} />
        <MessageForm
          onSubmit={msg => { msg && this.props.store.thread.addComment(msg) }} />
      </Comment.Group>
    )
  }
  renderItem (comment) {
    return (
      <Comment key={comment.id}>
        <Comment.Avatar as={Image} circular src={comment.image} />
        <Comment.Content>
          <Comment.Author as='a'>{comment.username}</Comment.Author>
          <Comment.Metadata>
            <Moment fromNow>{comment.date}</Moment>
          </Comment.Metadata>
          <Comment.Text>{comment.body || comment.extraText || comment.summary}</Comment.Text>
          {/* <Comment.Actions>
            <a>Reply</a>
          </Comment.Actions> */}
        </Comment.Content>
      </Comment>
    )
  }
}

export default CommentsList
