import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Card, Icon, Image } from 'semantic-ui-react'
import Moment from 'react-moment'

@inject('store') @observer
class SettingsCard extends Component {
  render () {
    const { store } = this.props
    return (
      <Card style={{ width: '100%' }}>
        { store.peer.profile.avatar && <Image src={store.peer.profile.avatar} /> }
        <Card.Content>
          <Card.Header>Settings</Card.Header>
          <Card.Meta>
            Updated <Moment fromNow>{store.peer.profile.updated}</Moment>
          </Card.Meta>
          <Card.Description>Peer Id: {store.peer.profile.id}</Card.Description>
        </Card.Content>
        <Card.Content extra>
          <Icon name='user' />{'??'} Peers
        </Card.Content>
      </Card>
    )
  }
}

export default SettingsCard
