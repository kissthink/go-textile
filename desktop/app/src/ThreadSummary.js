import React, { Component } from 'react'
import { Header, Icon, List } from 'semantic-ui-react'
import { observer, inject } from 'mobx-react'
import Moment from 'react-moment'

@inject('store') @observer
class ThreadSummary extends Component {
  render () {
    const { currentThread, threads } = this.props.store.peer
    if (Object.keys(threads).length < 1 || !currentThread) {
      return null
    }
    const info = threads[currentThread]
    return (
      <div>
        <Header as='h3' style={{ margin: 0 }}>{info.name}</Header>
        <List bulleted horizontal style={{ color: 'gray', fontSize: '0.8em' }}>
          <List.Item>
            <Icon name='user outline' />{info.peer_cnt}
          </List.Item>
          <List.Item>
            <Icon name='file outline' />{info.file_cnt}
          </List.Item>
          <List.Item>
            <Icon name='block layout' />{info.block_cnt}
          </List.Item>
          <List.Item>
            <Icon name='comment outline' />{info.block_cnt}
          </List.Item>
          <List.Item>
            <Icon name='heart outline' />{info.block_cnt}
          </List.Item>
          <List.Item>
            Updated <Moment fromNow>{info.head.date}</Moment>
          </List.Item>
        </List>
      </div>
    )
  }
}

export default ThreadSummary
