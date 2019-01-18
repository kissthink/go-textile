import { observable, computed, action } from 'mobx'
import { API } from './TextileAPI'

export class ChatStore {
    @observable thread = '12D3KooWQLTj6AZVo6pbL7rgvY9WY1KRGusraEQ2UnpjuPq9FYEb'
    @computed get messages () {
      return API.getMessages({ thread: this.thread, offset: '', limit: 20 })
        .then(list => list)
    }
    @action addMessage (body) {
      API.addMessage(body, this.thread)
        .then(message => this.messages.push(message))
    }
}

export default class MainStore {
    @observable isLoading = false
    @observable chat = new ChatStore({})
    @computed get username () {
      return !this.isLoading ? API.getProfile().then(profile => profile.username) : null
    }
    @observable threads = []
    @action updateThreads () {
      API.getThreads()
        .then(list => {
          this.threads = list
        })
    }
}
