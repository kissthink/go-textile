import React, { Component } from 'react'
import { Menu, Header } from 'semantic-ui-react'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class ThreadList extends Component {
  componentDidMount () {
    console.log('thread list')
    const { store } = this.props
    if (store.threads.length > 0 && store.thread.info === null) {
      store.thread.setThread(store.threads[0])
    }
  }
  render () {
    const { store } = this.props
    return (
      <Menu fluid vertical tabular>
        {/* <Header as='h3' dividing>
          Threads
        </Header> */}
        {store.threads && store.threads.slice().map(thread => this.renderItem(thread))}
      </Menu>
    )
  }
  renderItem (thread) {
    const { store } = this.props
    return (
      <Menu.Item
        key={thread.id}
        onClick={() => {
          store.thread.setThread(thread)
        }}
        active={store.thread.info && thread.id === store.thread.info.id}
      >
        <Header as='h3' style={{ margin: 0 }}>{thread.name}</Header>
      </Menu.Item>
    )
  }
}

export default ThreadList
