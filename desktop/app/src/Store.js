import { observable, computed, action, reaction } from 'mobx'
import { API } from './TextileAPI'
import { basicInfo, Blocks } from './TextileAPI/helpers'

// don't allow state modifications outside actions
// configure({ enforceActions: 'strict' })

export class ThreadStore {
  @observable info = null
  @observable updates = []
  @observable currentUpdate = null
  constructor () {
    reaction(() => this.info, newInfo => {
      this.getUpdates()
    })
  }
  @action getUpdates (limit) {
    API.getBlocks({ thread: this.info.id, offset: '', limit: limit || 50 })
      .then(Blocks.processBlocks)
      .then(list => { this.updates.replace(list) })
  }
  @action addMessage (body) {
    API.addMessage(body, this.info.id)
      .then(async message => {
        this.updates.unshift(await Blocks.processBlock(message))
      })
  }
  @action addComment (body) {
    API.addComment(body, this.currentUpdate.id)
      .then(async message => {
        const res = basicInfo(message)
        res.body = message.body
        this.updates
          .filter(u => u && u.id === this.currentUpdate.id)[0].comments.unshift(res)
      })
  }
  @action createInvite () {
    return API.createPublicInvite({ opts: { thread: this.info.id } })
  }
  @action setThread (info) {
    this.info = info
  }
  @action setCurrentUpdate (update) {
    this.currentUpdate = update
  }
  @action clearCurrentUpdate () {
    this.currentUpdate = null
  }
  @computed get peers () {
    if (!this.update) {
      return {}
    }
    return this.updates.reduce((a, b) => {
      const user = {}
      user[b.author_id] = {
        url: b.image,
        username: b.username,
        date: b.date
      }
      return { ...a, ...user }
    }, {})
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
  @observable status = 'init' // initializing, starting, ready
  @observable threads = []
  @observable username = null
  @observable thread = new ThreadStore()
  @observable ui = new UIStore()
  constructor () {
    reaction(() => this.isLoading, (isLoading, reaction) => {
      if (!isLoading) {
        this.getUsername()
        this.getThreads()
        reaction.dispose()
      }
    })
  }
  @action getUsername () {
    API.getProfile().then(p => {
      this.username = p.username || p.address.slice(-8)
    })
  }
  @action getThreads () {
    API.getThreads().then(list => {
      this.threads.replace(list)
    })
  }
  @action
  setStatus (status) {
    this.status = status
  }
  @computed
  get isLoading () {
    return this.status !== 'ready'
  }
}
