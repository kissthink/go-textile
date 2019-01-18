import React, { Component } from 'react'
import { Button, TextInput } from 'react-desktop/macOs'
import { observer } from 'mobx-react'

export default @observer class SendMessageForm extends Component {
  onSubmit (event) {
    const { store } = this.props
    event.preventDefault()
    store.addMessage(event.target.value)
  }

  render () {
    // const { store } = this.props
    return (
      <div className='send-message-form-container'>
        <form onSubmit={this.onSubmit} className='send-message-form'>
          <TextInput
            type='text'
            // value={state.chat.message}
            className='message-input'
          />
          <Button color='blue' type='submit'>
            Send
          </Button>
        </form>
      </div>
    )
  }
}
