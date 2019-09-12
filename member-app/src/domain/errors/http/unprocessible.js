const HTTPError = require('./http-error');

function UnprocessibleEntityErr(message) {
  this.httpErr = new HTTPError(422, message);
}

UnprocessibleEntityErr.send = function(res) {
  this.httpErr.send(res);
};

module.exports = UnprocessibleEntityErr;
