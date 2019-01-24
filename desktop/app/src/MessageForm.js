import React, { Component } from 'react'
import { Form } from 'semantic-ui-react'
import { observer, inject } from 'mobx-react'

@inject('store') @observer
class MessageForm extends Component {
  constructor (props) {
    super(props)
    this.state = { message: '' }
  }
  handleSubmit = () => {
    this.props.store.thread.addMessage(this.state.message)
    this.setState({ message: '' })
  }
  // use setState here... not propagating `message` to store
  handleChange = (e, { name, value }) => this.setState({ [name]: value })
  render () {
    const { message } = this.state
    return (
      <Form onSubmit={this.handleSubmit}>
        <Form.Group>
          <Form.Input
            placeholder='Type something...'
            name='message'
            value={message}
            width={14}
            onChange={this.handleChange} />
          <Form.Button
            content='Send'
          />
        </Form.Group>
      </Form>
    )
  }
}

export default MessageForm
