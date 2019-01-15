/* global asticode astilectron */
import React, { Component } from 'react'
import chat, Chat from './Chat'
import main, Main from './Main'

export const app = {
  initialState: () => Object.assign(
    main.initialState(),
    chat.initialState()
  ),
  actions: update => Object.assign({},
    main.actions(update),
    chat.actions(update),
  )
}

export class App extends Component {
  constructor (props) {
    super(props)
    this.state = props.states()
  }

  componentDidMount() {
    this.props.states.map(state => {
      return this.setState(state)
    })
    const { actions } = this.props
    astilectron.onMessage(message => {
      console.log(message)
      if (message.name === 'status') {
        asticode.notifier.info("Status: " + message.status)
        actions.setStatus(message.status)
        // TODO: This doesn't really belong here
        if (message.status === 'ready') {
          astilectron.sendMessage({ name: 'username' }, message => {
            if (message.name === 'username.callback') {
              asticode.notifier.info("Username: " + message.payload.username)
            }
            actions.setUsername(message.payload.username),
            actions.setScreen('chat')
          })
        }
      }
    })
  }

  render () {
    const state = this.state
    const { actions } = this.props
    switch (this.state.screen) {
      case 'chat':
        return <Chat state={state} actions={actions} />
      default:
        return <Main />
    }
  }
}
