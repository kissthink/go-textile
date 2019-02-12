import React, { Component } from 'react'
import { Menu, Header } from 'semantic-ui-react'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class ThreadList extends Component {
  componentDidMount () {
    const { store } = this.props
    if (Object.keys(store.peer.threads).length > 0 && store.peer.currentThread === null) {
      store.peer.currentThread = Object.keys(store.peer.threads)[0]
    }
  }
  render () {
    const { store } = this.props
    return (
      <Menu fluid vertical tabular>
        {/* <Header as='h3' dividing>
          Threads
        </Header> */}
        {store.peer.threads && Object
          .entries(store.peer.threads)
          .map(([id, thread]) => this.renderItem(id, thread))}
      </Menu>
    )
  }
  renderItem (id, thread) {
    const { store } = this.props
    return (
      <Menu.Item
        key={id}
        name={id}
        onClick={() => {
          store.peer.currentThread = id
        }}
        active={store.peer.currentThread && thread.id === store.peer.currentThread}
      >
        <Header as='h3' style={{ margin: 0 }}>{thread.name}</Header>
      </Menu.Item>
    )
  }
}

export default ThreadList
