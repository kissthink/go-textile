import React, { Component } from 'react'
import {
  ListView,
  ListViewSection,
  ListViewRow,
  Text
} from 'react-desktop/macOs'

class ThreadList extends Component {
  render () {
    return (
      <ListView className='online-list'>
        <ListViewSection>
          {this.props.threads && this.props.threads.map((thread, index) => {
            return this.renderItem(thread.name, thread.id)
          })}
        </ListViewSection>
      </ListView>
    )
  }

  renderItem (name, id) {
    return (
      <ListViewRow key={id}>
        <div
          className='online-list-item'
          style={{
            background: id === this.props.thread_id ? '#6BD761' : 'gray'
          }}
        />
        <Text color='#414141' size='13'>
          {name}{' '}
        </Text>{' '}
      </ListViewRow>
    )
  }
}

export default ThreadList
