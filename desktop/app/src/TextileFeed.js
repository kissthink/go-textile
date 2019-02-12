import React, { Component } from 'react'
import MessageForm from './MessageForm'
import ThreadList from './ThreadList'
import { observer, inject } from 'mobx-react'
import ThreadFeed from './ThreadFeed'
import { Segment, Grid } from 'semantic-ui-react'
import ThreadSummary from './ThreadSummary'

@inject('store') @observer
class TextileFeed extends Component {
  render () {
    return (
      <Grid style={{ height: '100vh' }} columns={2}>
        <Grid.Row stretched style={{ paddingBottom: '0.5rem' }}>
          <Grid.Column width={3}>
            <ThreadList />
          </Grid.Column>
          <Grid.Column width={13}>
            <ThreadSummary />
            <Segment basic style={{ margin: 0, height: 'calc(100% - 100px)' }}>
              <ThreadFeed />
            </Segment>
            <MessageForm onSubmit={msg => { msg && this.props.store.thread.addMessage(msg) }} />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    )
  }
}

export default TextileFeed
