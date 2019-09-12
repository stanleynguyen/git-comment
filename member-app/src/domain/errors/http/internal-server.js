const HTTPError = require('./http-error');

function InternalServerErr(message) {
  this.httpErr = new HTTPError(500, message);
}

InternalServerErr.send = function(res) {
  this.httpErr.send(res);
};

module.exports = InternalServerErr;
