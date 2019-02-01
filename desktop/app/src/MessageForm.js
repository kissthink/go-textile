import React, { Component } from 'react'
import { Form } from 'semantic-ui-react'

class MessageForm extends Component {
  constructor (props) {
    super(props)
    this.state = { message: '' }
  }
  handleSubmit = () => {
    this.props.onSubmit(this.state.message)
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
            // width={15}
            fluid
            action={{ icon: 'search' }}
            onChange={this.handleChange} />
          {/* <Form.Button icon='send' /> */}
        </Form.Group>
      </Form>
    )
  }
}

export default MessageForm
