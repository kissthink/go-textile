import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Sidebar, Segment, Icon } from 'semantic-ui-react'
import CommentsList from './CommentsList'

@inject('store') @observer
class InfoSidebar extends Component {
  render () {
    const { store } = this.props
    const info = (kind => {
      switch (kind) {
        case 'likes':
          return <CommentsList />
        case 'comments':
          return <CommentsList />
        default:
          return null
      }
    })(store.ui.infoType)
    return (
      <Sidebar
        as={Segment}
        animation='overlay'
        icon='labeled'
        width='wide'
        direction='right'
        visible={store.ui.infoType !== null}
      >
        <Icon name='close' style={{ float: 'right' }} onClick={() => {
          store.ui.toggleInfo(null)
        }} />
        {info}
      </Sidebar>
    )
  }
}

export default InfoSidebar
