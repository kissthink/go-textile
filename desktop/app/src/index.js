/* global asticode */
import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import {app, App} from './App'
import * as serviceWorker from './serviceWorker'
import flyd from 'flyd'

// Meiosis-based state management
const update = flyd.stream()
const states = flyd.scan((state, patch) => patch(state),
  app.initialState(), update)
const actions = app.actions(update)

document.addEventListener('astilectron-ready', function () {
  ReactDOM.render(<App states={states} actions={actions} />, document.getElementById('root'))
  // If you want your app to work offline and load faster, you can change
  // unregister() to register() below. Note this comes with some pitfalls.
  // Learn more about service workers: http://bit.ly/CRA-PWA
  serviceWorker.unregister()
  asticode.loader.init()
  asticode.modaler.init()
  asticode.notifier.init()
})
