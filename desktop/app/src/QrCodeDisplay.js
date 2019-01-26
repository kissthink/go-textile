import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Message, Segment } from 'semantic-ui-react'
import QRCode from 'qrcode.react'

const createUrl = async (store) => {
  const inviter = store.username
  const name = store.thread.info.name
  const invite = await store.thread.createInvite()
  return `https://www.textile.photos/invites/new#id=${invite.id}&key=${invite.key}&inviter=${inviter}&name=${encodeURI(name)}`
}

@inject('store') @observer
class QrCodeReader extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }
  componentDidMount () {
    createUrl(this.props.store).then(url => this.setState({ url }))
  }
  render () {
    // const { store } = this.props
    return (
      <Segment style={{ width: '50%', maxWidth: '450px', margin: 'auto' }}>
        <Message
          header='Invite your peers!'
          content={'Point your phone\'s camera at the Qr Code to join the thread.'}
        />
        <Segment vertical style={{ padding: 0 }} textAlign='center'>
          {this.state.url && <QRCode size={300} value={this.state.url} />}
        </Segment>
      </Segment>
    )
  }
}

export default QrCodeReader
