
export const fetchWrapper = {
  get,
  post,
  put,
  delete: _delete,
}
const baseUrl = process.env.VUE_APP_ROOT_API;
class ResponseModel {
  constructor(status, body) {
    this.Status = status
    this.Body = body
  }
}

function get(url) {
  const requestOptions = {
    method: 'GET',
  }
  url = `${baseUrl}/${url}`
  return fetch(url, requestOptions).then(handleResponse)
}

function post(url, body) {
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");
  var urlencoded = new URLSearchParams();
  for (const [key, value] of Object.entries(body)) {
    urlencoded.append(key, value);
  }
  const requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: urlencoded,
  }
  url = `${baseUrl}/${url}`
  return fetch(url, requestOptions).then(handleResponse)
}

function put(url, body) {
  const requestOptions = {
    method: 'PUT',
    body: JSON.stringify(body),
  }
  url = `${baseUrl}/${url}`
  return fetch(url, requestOptions).then(handleResponse)
}

// prefixed with underscored because delete is a reserved word in javascript
function _delete(url) {
  const requestOptions = {
    method: 'DELETE',
  }
  url = `${baseUrl}/${url}`
  return fetch(url, requestOptions).then(handleResponse)
}

// helper functions

/**
 * Returns the sum of a and b
 * @returns {ResponseModel} includes Status as integer and Body as json model or text
 */
function handleResponse(response) {
  return response.text().then((text) => {
    let data
    try {
      data = text && JSON.parse(text)
    } catch (e) {
      data = text
    }
    return new ResponseModel(response.status, data)
    // return data
  })
}