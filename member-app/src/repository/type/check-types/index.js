const checkTypes = require('check-types');

function CheckTypes(value) {
  this.value = value;
  this.validated = true;
}
CheckTypes.check = function(value) {
  return new CheckTypes(value);
};
CheckTypes.prototype.isString = function() {
  this.validated = this.validated && checkTypes.maybe.string(this.value);
  return this;
};
CheckTypes.prototype.isNumber = function() {
  this.validated = this.validated && checkTypes.maybe.number(this.value);
  return this;
};
CheckTypes.prototype.isEnum = function(choices) {
  let ok = true;
  if (this.value) {
    ok = choices.includes(this.value);
  }
  this.validated = this.validated && ok;
  return this;
};
CheckTypes.prototype.isRequired = function() {
  this.validated =
    this.validated &&
    checkTypes.not.null(this.value) &&
    checkTypes.not.undefined(this.value);
  return this;
};
CheckTypes.prototype.validate = function() {
  return this.validated;
};

module.exports = CheckTypes;
