import React, { Component } from 'react'
import { List, Image } from 'semantic-ui-react'
import Moment from 'react-moment'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class LikesList extends Component {
  render () {
    const { store } = this.props
    return (
      <List>
        {store.thread.currentUpdate && store.thread.currentUpdate.likes
          .slice()
          .map(like => this.renderItem(like))}
      </List>
    )
  }
  renderItem (like) {
    return (
      <List.Item key={like.id}>
        <Image avatar src={like.image} />
        <List.Content>
          <List.Header as='a'>{like.username}</List.Header>
          <List.Description>
            Liked this <Moment fromNow>{like.date}</Moment>
          </List.Description>
        </List.Content>
      </List.Item>
    )
  }
}

export default LikesList
