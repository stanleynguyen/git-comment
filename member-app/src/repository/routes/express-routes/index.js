const express = require('express');

function ExpressRoutes(memberCtrl) {
  this.router = express.Router();
  this.memberCtrl = memberCtrl;
}
ExpressRoutes.prototype.initializeAPIRoutes = function() {
  this.router.get(
    '/orgs/:org/members',
    this.memberCtrl.getMembersOfOrg.bind(this.memberCtrl),
  );

  return this;
};
ExpressRoutes.prototype.getRouter = function() {
  return this.router;
};

module.exports = ExpressRoutes;
