const { handleGenericHTTPErr } = require('../helpers');
const { UnprocessibleEntityErr } = require('../../../domain/errors/http');

function MembersController(membersUsecase, typeChecker) {
  this.membersUsecase = membersUsecase;
  this.typeChecker = typeChecker;
}

MembersController.prototype.getMembersOfOrg = async function(req, res) {
  try {
    const { org } = req.params;
    const { page } = req.query;
    if (
      !this.typeChecker
        .check(org)
        .isString()
        .isRequired()
        .validate()
    ) {
      throw new UnprocessibleEntityErr('Invalid org handle');
    }
    if (
      !this.typeChecker
        .check(page)
        .isString()
        .validate()
    ) {
      throw new UnprocessibleEntityErr('Invalid page cursor');
    }

    const resp = await this.membersUsecase.getMembers(org, page);
    res.status(200).json(resp);
  } catch (e) {
    handleGenericHTTPErr(e, res);
  }
};

module.exports = MembersController;
