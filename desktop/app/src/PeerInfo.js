import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Card, Icon, Image, Input, Form } from 'semantic-ui-react'
import Moment from 'react-moment'
import { GATEWAY_URL } from './TextileAPI'

@inject('store') @observer
class PeerInfoCard extends Component {
  handleSubmit = e => {
    e.preventDefault()
    this.props.store.peer.setUsername(this.inputRef.inputRef.value)
  }
  handleCopy = e => {
    const el = document.createElement('textarea')
    el.value = this.props.store.peer.profile.id
    el.setAttribute('readonly', '')
    el.style.position = 'absolute'
    el.style.left = '-9999px'
    document.body.appendChild(el)
    el.select()
    document.execCommand('copy')
    document.body.removeChild(el)
  }
  render () {
    const { store } = this.props
    return (
      <Card style={{ width: '100%' }}>
        <Image size='medium' circular src={store.peer.profile.avatar ? `${GATEWAY_URL}/ipfs/${store.peer.profile.avatar}/0/large/d` : 'https://react.semantic-ui.com/images/wireframe/image.png'}
          label={{ as: 'a', corner: 'left', icon: 'edit' }}
        />
        <Card.Content>
          <Card.Header>
            <Form style={{ fontSize: '1em' }} onSubmit={this.handleSubmit}>
              <Form.Field>
                <Icon link onClick={() => { this.inputRef.focus() }} name='pencil' size='small' style={{ paddingRight: '2.3em' }} />
                <Input style={{ width: 'initial' }} ref={c => { this.inputRef = c }} transparent defaultValue={store.peer.profile.username} />
              </Form.Field>
            </Form>
          </Card.Header>
          <Card.Description>
            {document.queryCommandSupported('copy') && <Icon link name='copy' onClick={this.handleCopy} style={{ paddingRight: '2.3em' }} />}
            {store.peer.profile.id}
          </Card.Description>
        </Card.Content>
        <Card.Content extra>
          Updated <Moment fromNow>{store.peer.profile.updated}</Moment>
        </Card.Content>
      </Card>
    )
  }
}

export default PeerInfoCard
