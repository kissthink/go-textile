/* global asticode */
import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import * as serviceWorker from './serviceWorker'
import MainStore from './Store'

document.addEventListener('astilectron-ready', () => {
  const store = new MainStore()
  window.store = store

  ReactDOM.render(<App store={store} />, document.getElementById('root'))
  // If you want your app to work offline and load faster, you can change
  // unregister() to register() below. Note this comes with some pitfalls.
  // Learn more about service workers: http://bit.ly/CRA-PWA
  serviceWorker.unregister()
  asticode.loader.init()
  asticode.modaler.init()
  asticode.notifier.init()
})
