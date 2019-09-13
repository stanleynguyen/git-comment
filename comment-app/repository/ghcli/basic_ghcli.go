package ghcli

import (
	"fmt"
	"github.com/stanleynguyen/git-comment/comment-app/repository"
	"net/http"

	"github.com/stanleynguyen/git-comment/comment-app/domain"
)

type basicGithubCli struct{}

// OrgExists check if a Github org exists
func (c *basicGithubCli) OrgExists(org string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/orgs/%s", org))
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else if resp.StatusCode >= 300 {
		return false, domain.NewErrorInternalServer("Problems with Github service")
	}

	return true, nil
}

// NewBasicGithubClient generate new instance of basic Github API client
func NewBasicGithubClient() repository.GithubCli {
	return &basicGithubCli{}
}
