import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { Card, Input, Radio, Button, Icon } from 'semantic-ui-react'

@inject('store') @observer
class SettingsCard extends Component {
  constructor (props) {
    super(props)
    this.state = {
      backupEnabled: false,
      backupLocation: ''
    }
  }
  handleSubmit = () => {
    this.props.onSubmit(this.state.message)
    this.setState({ message: '' })
  }
  // use setState here... not propagating `message` to store
  handleCheckbox = (e, data) => {
    this.setState({ [data.name]: data.checked })
  }
  handleInput = (e, data) => {
    this.setState({ [data.name]: data.value })
  }
  render () {
    return (
      <Card style={{ width: '100%' }}>
        <Card.Content>
          <Card.Header>Settings</Card.Header>
          <Card.Description>
            <h3>
              <Radio
                toggle
                name='backupEnabled'
                checked={this.state.backupEnabled}
                label='Enable file backup?'
                onClick={this.handleCheckbox}
              />
            </h3>
            <Input
              type='text'
              disabled={!this.state.backupEnabled}
              fluid
              name='backupLocation'
              onChange={this.handleInput}
              onSubmit={this.props.store.setBackupLocation}
              value={this.state.backupLocation}
              action={
                <label for='folder-select'>
                  <Button icon attached='right'>
                    <Icon name='folder open outline' />
                    <input
                      disabled={!this.state.backupEnabled}
                      type='file'
                      id='folder-select'
                      style={{ display: 'none' }}
                      webkitdirectory='' directory=''
                      onChange={this.handleInput}
                    />
                  </Button>
                </label>
              }
              placeholder='Choose location...'
            />
          </Card.Description>
        </Card.Content>
      </Card>
    )
  }
}

export default SettingsCard
