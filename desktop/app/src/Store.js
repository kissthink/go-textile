import { observable, computed, action, reaction } from 'mobx'
import { API } from './TextileAPI'

// don't allow state modifications outside actions
// configure({ enforceActions: 'strict' })

export class PeerStore {
  @observable profile = {}
  @observable threads = {}
  @observable currentThread = null
  @action async getProfile () {
    const profile = await API.getProfile()
    profile.username = profile.username || profile.address.slice(-8)
    this.profile = profile
  }
  @action async setUsername (username) {
    await API.setUsername(username)
    this.profile.username = username
  }
  @action async updateThreads () {
    const threads = await API.getThreads()
    for (const thread of threads) {
      this.updateThread(thread.id)
    }
  }
  @action async updateThread (thread, limit) {
    // const newLimit = limit || (this.threads[thread] ? this.threads[thread].length : null) || 50
    // const list = await API.getFeed({ thread: thread, offset: '', limit: newLimit })
    // this.threads[thread].replace(list)
  }
  @action async addMessage (body, thread) {
    this.threads[thread].unshift(await API.addMessage(body, thread))
  }
  @action async addLike (block) {
    API.addLike(block)
    // TODO: add this to the appropriate thread.block
  }
  @action async addComment (body, block) {
    API.addComment(body, block)
    // TODO: add this to the appropriate thread.block
  }
  @action async addFlag (block) {
    API.addFlag(block)
    // TODO: add this to the appropriate thread.block
  }
}

export class UIStore {
  @observable screen = 'loading' // loading, feed, settings, qrcode
  @observable menuVisible = false
  @observable infoType = null
  @action setScreen (screen) {
    this.screen = screen
  }
  @action toggleMenu () {
    this.menuVisible = !this.menuVisible
  }
  @action toggleInfo (kind) {
    this.infoType = kind
  }
}

export default class MainStore {
  @observable status = 'initializing' // initializing, starting, ready
  @observable peer = new PeerStore()
  @observable ui = new UIStore()
  constructor () {
    reaction(() => this.isLoading, (isLoading, reaction) => {
      if (!isLoading) {
        this.peer.getProfile()
        this.peer.updateThreads()
      }
    })
  }
  @action setStatus (status) {
    this.status = status
  }
  @computed get isLoading () {
    return this.status !== 'ready'
  }
}
