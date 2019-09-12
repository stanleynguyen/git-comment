function MembersUsecase(graphCli) {
  this.graphCli = graphCli;
}

MembersUsecase.prototype.getMembers = async function(org, cursor) {
  const query = `{
    organization(login: "${org}") {
      membersWithRole(first: 100${cursor ? `, after: "${cursor}"` : ''}) {
        pageInfo {
          endCursor
        }
        edges {
          node {
            login
            avatarUrl
            followers {
              totalCount
            }
          }
        }
      }
    }
  }`;
  const resp = await this.graphCli.request(query);
  const membersData = resp.organization.membersWithRole;
  return {
    nextPageCursor: membersData.pageInfo.endCursor,
    members: membersData.edges
      .map(mapNode2ReturnObj)
      .sort(sortByNumberOFollowers),
  };
};

const mapNode2ReturnObj = v => ({
  login: v.node.login,
  avatarUrl: v.node.avatarUrl,
  followers: v.node.followers.totalCount,
});

const sortByNumberOFollowers = (a, b) => b.followers - a.followers;

module.exports = MembersUsecase;
