const fs = require('fs').promises

async function getFileResponse() {
  // no streaming support here. :(
  const contents = await fs.readFile(require.resolve('./resume.pdf'), 'binary')
  const body = Buffer.from(contents, 'binary').toString('base64')
  return {
    statusCode: 200,
    headers: { 'Content-Type': 'application/pdf' },
    body,
    isBase64Encoded: true
  }
}

function isAuthorized(event) {
  const authHeader = event.headers.authorization
  if (!authHeader) {
    return false
  }

  const [authScheme, encodedCreds] = authHeader.split(' ')
  if (authScheme !== 'Basic' || !encodedCreds) {
    return false
  }

  const [suppliedUsername, suppliedPassword] = Buffer.from(encodedCreds, 'base64').toString().split(':')
  if (!suppliedUsername || !suppliedPassword) {
    return false
  }

  const credentials = JSON.parse(process.env.RESUME_CREDENTIALS)
  return credentials.some(cred => suppliedUsername === cred.username && suppliedPassword === cred.password)
}


exports.handler = async function (event, context) {
  if (!isAuthorized(event)) {
    return {
      statusCode: 401,
      headers: { 'WWW-Authenticate': 'Basic realm="No credentials? Send me an email :)"' },
      body: 'Unauthorized'
    }
  }
  return getFileResponse()
}
