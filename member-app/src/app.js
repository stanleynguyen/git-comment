const express = require('express');
const GraphCli = require('./repository/graph-cli');
const { MembersUsecase } = require('./usecase');
const { MembersHTTPController } = require('./delivery/http');
const TypeRepo = require('./repository/type');
const RoutesRepo = require('./repository/routes');

module.exports = {
  inDevMode: function() {
    require('dotenv').config();
    this.app = this.bootstrap();
    return this;
  },

  inProdMode: function() {
    this.app = this.bootstrap();
    return this;
  },

  start: function() {
    this.app.listen(process.env.PORT, () => {
      console.log(`UP AND RUNNING ON ${process.env.PORT}`);
    });
  },

  bootstrap: function() {
    const graphCli = new GraphCli('request', 'https://api.github.com/graphql', {
      headers: {
        Authorization: `bearer ${process.env.GITHUB_TOKEN}`,
      },
    });
    const memberUsecase = new MembersUsecase(graphCli);
    const typeChecker = new TypeRepo('check-types');
    const memberCtrl = new MembersHTTPController(memberUsecase, typeChecker);
    const expressRoutesRepo = RoutesRepo('express');
    const routes = new expressRoutesRepo(memberCtrl);
    const app = express();
    app.use(routes.initializeAPIRoutes().getRouter());
    return app;
  },
};
