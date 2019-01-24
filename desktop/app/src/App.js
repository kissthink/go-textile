/* global astilectron, window */
import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import Root from './Root'
import './App.css'

@inject('store') @observer
class App extends Component {
  componentDidMount () {
    const { store } = this.props
    if ('astilectron' in window) {
      astilectron.onMessage(message => {
        switch (message.name) {
          case 'status':
            store.setStatus(message.status)
            store.setScreen('settings')
            break
          default:
            console.log(message)
        }
      })
    } else {
      setTimeout(() => {
        store.setStatus('starting')
      }, 500)
      setTimeout(() => {
        store.setStatus('ready')
        store.ui.setScreen('settings')
      }, 1000)
    }
  }

  render () {
    return (
      <div>
        <Root />
      </div>
    )
  }
}

export default App
