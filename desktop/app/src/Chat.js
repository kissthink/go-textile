/* global asticode astilectron */
import React, { Component } from 'react'
import MessageList from './MessageList'
import SendMessageForm from './SendMessageForm'
import ThreadList from './ThreadList'

const initialState = () => ({
  chat: {
    thread_id: '12D3KooWNVjmqQkSbdJ65JZBa8Tmw7uD6V3qJGqUyGrJBJQc4f7X',
    messages: [],
    threads: [],
    message: null,
  }
})

const actions = update => ({
  sendMessage: value =>
    update(state => set(state, ['chat', 'message'], value)),
  updateThreads: value =>
    update(state => set(state, ['chat', 'threads'], value)),
  updateMessages: value =>
    update(state => set(state, ['chat', 'messages'], value))
})

export const chat = {
  initialState,
  actions
}

class Chat extends Component {
  componentDidMount() {
    const { actions } = this.props
    astilectron.sendMessage({ name: 'threads' }, message => {
      if (message.name === 'threads.callback') {
        asticode.notifier.info('Threads: ' + message.payload.threads.length)
      }
      actions.updateThreads(message.payload.threads)
    })
  }

  onThreadSelect () {
    astilectron.sendMessage({
      name: 'messages',
      payload: this.state.thread_id
    }, message => {
      if (message.name === 'messages.callback') {
        asticode.notifier.info("Messages: " + message.payload.messages.length)
      }
      this.setState({
        messages: message.payload.messages,
      })
    })
  }

  onSend = body => {
    astilectron.sendMessage({
      name: 'message',
      payload: {
        thread_id: this.state.thread_id,
        body: body
      },
    }, message => {
      if (message.name !== 'error') {
        asticode.notifier.info(message)
      }
      this.setState({
        messages: [...this.state.messages, message.payload]
      })
    })
  }

  onSwitch = thread_id => {
    this.setState({
      thread_id: thread_id
    })
  }

  render () {
    const { state, actions } = this.props
    return (
      <div className="wrapper">
        <div>
          <ThreadList
            thread_id={state.thread_id}
            threads={state.threads}
          />
        </div>
        <div className="chat">
          <MessageList messages={state.messages} />
          <SendMessageForm onSend={actions.sendMessage} />
        </div>
      </div>
    )
  }
}

export default Chat
