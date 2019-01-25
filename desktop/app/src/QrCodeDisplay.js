import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Message, Segment } from 'semantic-ui-react'
import QRCode from 'qrcode.react'

const createUrl = (store) => {
  const { id, key, name } = store.thread.info
  const inviter = store.username
  return `https://www.textile.photos/invites/new#id=${id}&key=${key}&inviter=${inviter}&name=${encodeURI(name)}`
}

@inject('store') @observer
class QrCodeReader extends Component {
  componentDidMount () {
    const canvas = this.refs.canvas
    canvas.height = canvas.width
  }
  render () {
    const { store } = this.props
    return (
      <Segment style={{ width: '50%', maxWidth: '450px', margin: 'auto' }}>
        <Message
          header='Invite your peers!'
          content={'Point your phone\'s camera at the Qr Code to join the thread.'}
        />
        <Segment vertical style={{ padding: 0 }} textAlign='center'>
          <QRCode style={{ width: '100%' }} ref='canvas' size={300} value={createUrl(store)} />
        </Segment>
      </Segment>
    )
  }
}

export default QrCodeReader
