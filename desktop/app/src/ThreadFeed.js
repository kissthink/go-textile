import React, { Component } from 'react'
import { Feed, Icon, Image } from 'semantic-ui-react'
import { observer, inject } from 'mobx-react'
import Moment from 'react-moment'

@inject('store') @observer
class ThreadFeed extends Component {
  componentDidMount () {
    console.log('thread feed')
  }
  render () {
    const { updates } = this.props.store.thread
    return (
      <Feed style={{ paddingRight: '1em', height: '100%', overflowY: 'auto' }}>
        {updates && updates
          .slice()
          .reverse()
          .map(update => {
            return this.renderItem(update)
          })}
      </Feed>
    )
  }
  renderItem (item) {
    const { store } = this.props
    return (
      <Feed.Event key={item.id} onClick={() => {
        store.thread.setCurrentUpdate(item)
      }}>
        <Feed.Label>
          <Image avatar onError={i => { i.target.src = '' }} src={item.image} />
        </Feed.Label>
        <Feed.Content>
          <Feed.Summary>
            <Feed.User>{item.username}</Feed.User> {item.summary}
            <Feed.Date>
              <Moment fromNow>{item.date}</Moment>
            </Feed.Date>
            <Feed.Meta style={{ float: 'right' }}>
              <Feed.Like onClick={() => {
                store.ui.toggleInfo('likes')
              }}>
                <Icon name='like' />
                {item.likes && item.likes.length}
              </Feed.Like>
              <Feed.Like onClick={() => {
                store.ui.toggleInfo('comments')
              }}>
                <Icon color='green' name='comment' />
                {item.comments && item.comments.length}
              </Feed.Like>
            </Feed.Meta>
          </Feed.Summary>
          <Feed.Extra text>
            {item.extraText}
          </Feed.Extra>
          <Feed.Extra images>
            {item.extraImages && item.extraImages.map((img, i) => {
              return <a key={i} href={img}><img key={i} src={img} alt='posted' /></a>
            })}
          </Feed.Extra>
        </Feed.Content>
      </Feed.Event>
    )
  }
}

export default ThreadFeed
