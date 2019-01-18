import React, { Component } from 'react'
import MessageList from './MessageList'
import SendMessageForm from './SendMessageForm'
import ThreadList from './ThreadList'
import { observer } from 'mobx-react'

export default @observer class Chat extends Component {
  componentDidMount () {
    const { store } = this.props
    store.updateThreads()
  }
  render () {
    console.log('chat render')
    const { store } = this.props
    return (
      <div className='wrapper'>
        <div>
          <ThreadList store={store} />
        </div>
        <div className='chat'>
          <MessageList store={store.chat} />
          <SendMessageForm store={store.chat} />
        </div>
      </div>
    )
  }
}
