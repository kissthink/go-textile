/* global asticode astilectron */
import React, { Component } from 'react'
import { observer } from 'mobx-react'
import Root from './Root'

export @observer class App extends Component {
  componentDidMount () {
    const store = this.props.store
    astilectron.onMessage(message => {
      console.log(message)
      if (message.name === 'status') {
        asticode.notifier.info('Status: ' + message.status)
        store.status = message.status
        // TODO: This doesn't really belong here
        if (message.status === 'ready') {
          store.screen = 'chat'
        }
      }
    })
  }

  render () {
    console.log('app render')
    const store = this.props.store
    return <Root store={store} />
  }
}
