import React, { Component } from 'react'
import { ListView, ListViewSection, ListViewRow, Text } from 'react-desktop/macOs'
import { observer } from 'mobx-react'

export default @observer class MessageList extends Component {
  render () {
    const { store } = this.props
    return (
      <ListView>
        <ListViewSection>
          {[...store.messages].reverse().map(message =>
            this.renderItem(message)
          )}
        </ListViewSection>
      </ListView>
    )
  }
  renderItem (message) {
    return (
      <ListViewRow key={message.date} >
        <Text color='#414141' size='13' bold>
          {message.username}:
        </Text>
        <Text color='#414141' size='13'>
          {message.body}
        </Text>
      </ListViewRow>
    )
  }
}
