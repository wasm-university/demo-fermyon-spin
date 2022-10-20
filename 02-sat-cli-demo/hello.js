const http = require('http')

const requestListener = function (req, res) {
  console.log("👋 request from wasm module")
  res.writeHead(200);
  res.end("👋 Hello! 😀🎃")
}

console.log("🌍 listening on 9090")
const server = http.createServer(requestListener)
server.listen(9090)
