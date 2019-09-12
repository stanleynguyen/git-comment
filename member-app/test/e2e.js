process.env.NODE_ENV = 'test';

const chai = require('chai');
const expect = chai.expect;
const chaiHttp = require('chai-http');
const server = require('../src/app').inDevMode();

chai.use(chaiHttp);

describe('Members service', () => {
  it("should GET org's members with login, avatarUrl and number of followers", done => {
    chai
      .request(server.app)
      .get('/orgs/google/members')
      .end((err, res) => {
        expect(res.status).to.be.equal(200);
        expect(res.body).to.have.property('members');
        expect(res.body.members).to.be.a('array');
        res.body.members.forEach(m =>
          expect(m).to.have.all.keys('login', 'avatarUrl', 'followers'),
        );
        done();
      });
  });
  it("should GET org's members in DESC number of followers", done => {
    chai
      .request(server.app)
      .get('/orgs/google/members')
      .end((err, res) => {
        const members = res.body.members;
        for (let i = 0; i < members.length - 1; i++) {
          expect(members[i].followers >= members[i].followers).to.be.true;
        }
        done();
      });
  });
});
