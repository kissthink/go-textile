/* global fetch */
const API_ROOT = 'http://localhost:38100/api/v0'

const request = async (method, url, { args, opts, ctype, body }) => {
  let headers = {}
  if (args && args.length > 0) {
    headers['X-Textile-Args'] = opts.args.map(encodeURI).join(',')
  }
  if (opts && opts.length > 0) {
    headers['X-Textile-Opts'] = Object
      .entries(opts)
      .map((k, v) => k + '=' + encodeURI(v))
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
  getThreads: () => request('get', '/threads', {}),
  addMessage: (body, thread) => request('post', `/threads/${thread}/messages`, { args: [body] }),
  getProfile: () => request('get', '/profile', {})
}
