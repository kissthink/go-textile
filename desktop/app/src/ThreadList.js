import React, { Component } from 'react'
import { ListView, ListViewSection, ListViewRow, Text } from 'react-desktop/macOs'
import { observer } from 'mobx-react'

export default @observer class ThreadList extends Component {
  render () {
    const { store } = this.props
    return (
      <ListView className='online-list'>
        <ListViewSection>
          {store.chat.threads && store.chat.threads.map(thread => {
            return this.renderItem(thread)
          })}
        </ListViewSection>
      </ListView>
    )
  }

  renderItem (thread) {
    const { store } = this.props
    return (
      <ListViewRow key={thread.id}
        onClick={() => store.chat.setThread(thread.id)}
      >
        <div
          className='online-list-item'
          style={{
            background: thread.id === store.chat.thread ? '#6BD761' : 'gray'
          }}
        />
        <Text color='#414141' size='13'>
          {thread.name}{' '}
        </Text>{' '}
      </ListViewRow>
    )
  }
}
