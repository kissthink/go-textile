import React, { Component } from 'react'
import logo from './assets/img/logo.svg'
import { observer } from 'mobx-react'

export default @observer class Splash extends Component {
  render () {
    console.log('splash render')
    const store = this.props.store
    return (
      <div>
        <h1>
          {store.status}
        </h1>
        <img alt='' src={logo} width='200' />
      </div>
    )
  }
}
