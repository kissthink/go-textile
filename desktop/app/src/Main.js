import React, { Component } from 'react'
import logo from './assets/img/logo.svg'
import set from 'lodash.set'

const initialState = () => ({
  main: {
    username: null,
    status: 'loading'
  }
})

const actions = update => ({
  setStatus: value =>
    update(state => set(state, ['main', 'username'], value)),

  setUsername: value =>
    update(state => set(state, ['main', 'status'], value))
})

export const main = {
  initialState,
  actions
}

class Main extends Component {
  constructor () {
    super()
    this.state = {
      status: 'loading'
    }
  }

  render () {
    const { state } = this.props
    return (
      <div>
        <h1>
          Welcome {state.username}!
        </h1>
        <img alt='' src={logo} width='200' />
      </div>
    )
  }
}

export default Main
