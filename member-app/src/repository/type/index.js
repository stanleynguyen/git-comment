const CheckTypes = require('./check-types');

module.exports = function TypeChecker(lib) {
  switch (lib) {
    case 'check-types':
    default:
      return CheckTypes;
  }
};
