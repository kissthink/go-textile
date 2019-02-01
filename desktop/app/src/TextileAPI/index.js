/* global fetch */
export const API_ROOT = 'http://127.0.0.1:17103/api/v0'
export const GATEWAY_URL = 'http://127.0.0.1:32950'

const request = async (method, url, { args, opts, ctype, body }) => {
  let headers = {}
  if (args && args.length > 0) {
    headers['X-Textile-Args'] = args.map(encodeURI).join(',')
  }
  if (opts && Object.keys(opts).length > 0) {
    headers['X-Textile-Opts'] = Object
      .entries(opts)
      .map(([k, v]) => k + '=' + encodeURI(v))
      .join(',')
  }
  if (ctype) {
    headers['Content-Type'] = ctype
  }
  const response = await fetch(API_ROOT + url, { method, headers, body })
  return response.json()
}

export const API = {
  // getFile: (block, opts) => request('get', `/files/${block}`, opts),
  // getPeers: thread => request('get', `/threads/${thread}/peers`, {}),
  // getFiles: opts => request('get', '/files', opts),
  getComments: block => request('get', `/blocks/${block}/comments`, {}),
  addComment: (body, block) => request('post', `/blocks/${block}/comments`, { args: [body] }),
  getLikes: block => request('get', `/blocks/${block}/likes`, {}),
  addLike: block => request('post', `/blocks/${block}/likes`, {}),
  // getBlocks: opts => request('get', '/blocks', { opts }),
  getThreads: () => request('get', '/threads', {}),
  getFeed: opts => request('get', '/feed', { opts }),
  // getMessages: opts => request('get', '/messages', opts),
  addMessage: (body, thread) => request('post', `/threads/${thread}/messages`, { args: [body] }),
  getProfile: () => request('get', '/profile', {}),
  setUsername: name => request('post', '/profile/username', { args: [name] }),
  setAvatar: hash => request('post', '/profile/username', { args: [hash] }),
  checkMessages: () => request('post', '/cafes/messages', {})
  // joinPublicInvite: opts => request('post', '/invites', opts),
  // createPublicInvite: opts => request('post', '/invites', opts)
}
