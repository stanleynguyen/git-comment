const ExpressRoutes = require('./express-routes');

function Routes(lib) {
  switch (lib) {
    case 'express':
    default:
      return ExpressRoutes;
  }
}

module.exports = Routes;
