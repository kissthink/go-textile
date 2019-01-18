import React, { Component } from 'react'
import Splash from './Splash'
import Chat from './Chat'
import { observer } from 'mobx-react'

export default @observer class Root extends Component {
  render () {
    console.log('root render')
    const { store } = this.props
    switch (store.screen) {
      case 'chat':
        return <Chat store={store} />
      default:
        return <Splash store={store} />
    }
  }
}
