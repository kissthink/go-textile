import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import * as serviceWorker from './serviceWorker'
import MainStore from './Store'
import App from './App'
import { Provider } from 'mobx-react'
import 'semantic-ui-css/semantic.min.css'

const store = new MainStore()

const app = (
  <Provider store={store}>
    <App />
  </Provider>
)

// // document.addEventListener('astilectron-ready', () => {
ReactDOM.render(app, document.getElementById('root'))
// // If you want your app to work offline and load faster, you can change
// // unregister() to register() below. Note this comes with some pitfalls.
// // Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister()
// // })
