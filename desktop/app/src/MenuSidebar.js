import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Icon, Menu, Sidebar } from 'semantic-ui-react'

@inject('store') @observer
class MenuSidebar extends Component {
  handleItemClick = (e, { name }) => {
    this.props.store.ui.setScreen(name)
    this.props.store.ui.menuVisible = false
  }
  render () {
    const { store } = this.props
    return (
      <Sidebar
        as={Menu}
        animation='push'
        icon='labeled'
        vertical
        width='thin'
        visible={store.ui.menuVisible}
      >
        <Menu.Item as='a'
          onClick={this.handleItemClick}
          name='feed'
          active={store.ui.screen === 'feed'}
        >
          <Icon name='comments outline' />
            Feed
        </Menu.Item>
        {/* <Menu.Item as='a'
          onClick={this.handleItemClick}
          name='invite'
          active={store.ui.screen === 'invite'}
        >
          <Icon name='qrcode' />
          Invite
        </Menu.Item>
        <Menu.Item as='a'
          onClick={this.handleItemClick}
          name='join'
          active={store.ui.screen === 'join'}
        >
          <Icon name='qrcode' />
            Join
        </Menu.Item> */}
        <Menu.Item as='a'
          onClick={this.handleItemClick}
          name='settings'
          active={store.ui.screen === 'settings'}
        >
          <Icon name='settings' />
            Settings
        </Menu.Item>
      </Sidebar>
    )
  }
}

export default MenuSidebar
