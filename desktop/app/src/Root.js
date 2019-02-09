import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import MenuSidebar from './MenuSidebar'
import InfoSidebar from './InfoSidebar'
import TextileFeed from './TextileFeed'
// import QrCodeReader from './QrCodeReader'
// import QrCodeDisplay from './QrCodeDisplay'
import PeerInfo from './PeerInfo'
import Settings from './Settings'
import { Dimmer, Loader, Sidebar, Button, Segment } from 'semantic-ui-react'
@inject('store') @observer
class Root extends Component {
  render () {
    const { store } = this.props
    const view = (screen => {
      switch (screen) {
        case 'feed':
          return <TextileFeed />
        // case 'join':
        //   return <QrCodeReader />
        // case 'invite':
        //   return <QrCodeDisplay />
        case 'settings':
          return (
            <div>
              <Segment style={{ width: '80%', maxWidth: '500px', margin: 'auto' }}>
                <PeerInfo />
                <Settings />
              </Segment>
            </div>
          )
        default:
          return (
            <Dimmer inverted active={store.isLoading}>
              <Loader inverted size='massive'>{store.status}</Loader>
            </Dimmer>
          )
      }
    })(store.ui.screen)
    return (
      <div style={{ height: '100%', padding: '10px' }}>
        <Button
          onClick={() => store.ui.toggleMenu()}
          icon='bars'
          compact
          floated='left'
          toggle
          attached={store.ui.menuVisible ? 'left' : false}
          active={store.ui.menuVisible}
        />
        <Sidebar.Pushable>
          <MenuSidebar />
          <InfoSidebar />
          <Sidebar.Pusher>
            {view}
          </Sidebar.Pusher>
        </Sidebar.Pushable>
      </div>
    )
  }
}

export default Root
