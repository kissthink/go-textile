import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Card, Input, Radio, Button, Icon } from 'semantic-ui-react'

@inject('store') @observer
class SettingsCard extends Component {
  constructor (props) {
    super(props)
    this.state = {
      fileBackupEnabled: false,
      fileBackupLocation: ''
    }
  }
  handleChange = (event, data) => {
    if (data.checked !== undefined) {
      this.setState({ fileBackupEnabled: data.checked })
    }
  }
  render () {
    // const { store } = this.props
    return (
      <Card style={{ width: '100%' }}>
        <Card.Content>
          <Card.Header>Settings</Card.Header>
          <Card.Description>
            <h3>
              <Radio
                toggle
                label='Enable file backup?'
                onClick={this.handleChange}
              />
            </h3>
            <Input
              type='text'
              disabled={!this.state.fileBackupEnabled}
              fluid
              onChange={e => console.log(e)}
              action={
                <Button as='label' for='folder-select' attached='right'>
                  <Icon name='folder open outline' />
                  <input type='file' id='folder-select' style={{ display: 'none' }} />
                </Button>
              }
              placeholder='Choose location...'
              defaultValue=''
            />
          </Card.Description>
        </Card.Content>
      </Card>
    )
  }
}

export default SettingsCard
