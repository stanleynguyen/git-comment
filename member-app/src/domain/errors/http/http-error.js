function HTTPError(status, message) {
  this.status = status;
  this.message = message;
}
HTTPError.prototype.send = function(res) {
  res.status(this.status).json({ message: this.message });
};

module.exports = HTTPError;
