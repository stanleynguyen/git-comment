const { GraphQLClient } = require('graphql-request');

function GraphQLRequest(endpoint, options) {
  this.graphQLClient = new GraphQLClient(endpoint, options);
}

GraphQLRequest.prototype.request = async function(query) {
  return await this.graphQLClient.request(query);
};

module.exports = GraphQLRequest;
