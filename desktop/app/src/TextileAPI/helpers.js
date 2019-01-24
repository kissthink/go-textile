import { API, GATEWAY_URL } from './index'

const basicInfo = block => {
  return {
    id: block.id,
    image: `${GATEWAY_URL}/ipns/${block.author_id}/avatar/large`,
    date: block.date,
    username: block.username,
    author_id: block.author_id
  }
}

export const Blocks = {
  processBlocks: async blocks => {
    return Promise.all(blocks.map(Blocks.processBlock))
  },
  processBlock: async block => {
    let msg = basicInfo(block)
    switch (block.type) {
      case 'MESSAGE':
        msg.summary = 'said'
        msg.extraText = block.body
        break
      case 'JOIN':
        msg.summary = 'joined the group'
        break
      case 'FILES':
        const file = await API.getFile(block.id, {})
        msg.summary = 'added a photo'
        msg.extraText = file.caption
        msg.extraImages = [
          file.files.map(f => {
            let url
            if (f.links && f.links.large) {
              const hash = f.links.large.hash
              url = `${GATEWAY_URL}/ipfs/${hash}`
              const key = f.links.large.key
              if (key) {
                url += `?key=${key}`
              }
            }
            return url
          })
        ]
        msg.likes = file.likes ? file.likes.map(basicInfo) : []
        msg.comments = file.comments ? file.comments.map(comment => {
          return { ...basicInfo(comment), ...comment }
        }) : []
        return msg
      default:
        return
    }
    msg.comments = await API.getComments(block.id, {})
    msg.likes = await API.getLikes(block.id, {})
    return msg
  }
}
