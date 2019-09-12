const GraphQLRequest = require('./graphql-request');

function GraphCli(client = 'request', ...constructorArgs) {
  switch (client) {
    case 'request':
    default:
      this.client = new GraphQLRequest(...constructorArgs);
  }
}

GraphCli.prototype.request = async function(query) {
  return await this.client.request(query);
};

module.exports = GraphCli;
