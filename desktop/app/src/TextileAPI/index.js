/* global fetch */
export const API_ROOT = 'http://127.0.0.1:38100/api/v0'
export const GATEWAY_URL = 'http://127.0.0.1:3565'

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
  getMessages: opts => request('get', '/messages', opts),
  getFile: (block, opts) => request('get', `/files/${block}`, opts),
  getPeers: thread => request('get', `/threads/${thread}/peers`, {}),
  getFiles: opts => request('get', '/files', opts),
  getComments: block => request('get', `/blocks/${block}/comments`, {}),
  addComment: (body, block) => request('post', `/blocks/${block}/comments`, { args: [body] }),
  getLikes: block => request('get', `/blocks/${block}/likes`, {}),
  getBlocks: opts => request('get', '/blocks', { opts }),
  getThreads: () => request('get', '/threads', {}),
  addMessage: (body, thread) => request('post', `/threads/${thread}/messages`, { args: [body] }),
  getProfile: () => request('get', '/profile', {}),
  checkMessages: () => request('post', '/cafes/messages', {})
}
