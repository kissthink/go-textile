import React, { Component } from 'react'
import { Divider, Header, Comment, Form, Button } from 'semantic-ui-react'
import Moment from 'react-moment'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class CommentsList extends Component {
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
          .map(comment => this.renderItem(comment))}
        <Form reply>
          <Form.Input />
          <Button content='Reply' labelPosition='left' icon='edit' primary />
        </Form>
      </Comment.Group>
    )
  }
  renderItem (comment) {
    return (
      <Comment key={comment.id}>
        <Comment.Avatar as='a' src={comment.image} />
        <Comment.Content>
          <Comment.Author as='a'>{comment.username}</Comment.Author>
          <Comment.Metadata>
            <Moment fromNow>{comment.date}</Moment>
          </Comment.Metadata>
          <Comment.Text>{comment.body || comment.extraText}</Comment.Text>
          {/* <Comment.Actions>
            <a>Reply</a>
          </Comment.Actions> */}
        </Comment.Content>
      </Comment>
    )
  }
}

export default CommentsList
