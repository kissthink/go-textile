import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Button, Message, Icon, Segment } from 'semantic-ui-react'
import QrReader from 'react-qr-reader'
import parse from 'url-parse'
import * as qs from 'querystringify'

@inject('store') @observer
class QrCodeReader extends Component {
  constructor (props) {
    super(props)
    this.state = {
      result: 'Waiting...'
    }
  }
  handleScan = (data) => {
    if (data) {
      const invite = qs.parse(parse(data, true).hash)
      this.setState({
        result: `Would you like to join the '${invite.name}' group?`
      })
    }
  }
  handleError = (err) => {
    console.error(err)
  }
  render () {
    return (
      <Segment style={{ width: '50%', maxWidth: '450px', margin: 'auto' }}>
        <Message
          vertical
          attached='top'
          header='Join a group!'
          content='Point a Textile Photos invite Qr Code at the camera.'
        />
        <Segment vertical style={{ padding: 0 }}>
          <QrReader
            delay={this.state.delay}
            onError={this.handleError}
            onScan={this.handleScan}
            style={{ width: '100%', maxWidth: '500px' }}
          />
        </Segment>
        <Message attached='bottom' info vertical>
          <Icon name='help' />
          Ready to join '{this.state.result}'?
          <Button.Group>
            <Button>Cancel</Button>
            <Button.Or />
            <Button positive>Yes</Button>
          </Button.Group>
        </Message>
      </Segment>
    )
  }
}

export default QrCodeReader
