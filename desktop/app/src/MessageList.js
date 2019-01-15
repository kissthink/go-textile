import React, { Component } from 'react'
import {
  ListView,
  ListViewSection,
  ListViewRow,
  Text
} from 'react-desktop/macOs'

class MessageList extends Component {
  render () {
    return (
      <ListView>
        <ListViewSection>
          {[...this.props.messages].reverse().map((message, index) =>
            this.renderItem(message)
          )}
        </ListViewSection>
      </ListView>
    )
  }

  renderItem (message) {
    return (
      <ListViewRow key={message.date}>
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

export default MessageList
