package comment

import (
	"fmt"
	"github.com/stanleynguyen/git-comment/comment-app/domain"
	"github.com/stanleynguyen/git-comment/comment-app/repository"
	"github.com/stanleynguyen/git-comment/comment-app/usecase"
)

type commentUsecase struct {
	dbRepo repository.Persistence
	ghCli  repository.GithubCli
}

// NewCommentUsecase generate new comment usecase
func NewCommentUsecase(p repository.Persistence, c repository.GithubCli) usecase.CommentUsecase {
	return &commentUsecase{p, c}
}

// Create create and persist a comment against an existing Github org
func (u *commentUsecase) Create(org, comment string) (*domain.Comment, error) {
	exists, err := u.ghCli.OrgExists(org)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.NewErrorNotFound(fmt.Sprintf("Org %s not found", org))
	}

	newComment := &domain.Comment{Org: org, Comment: comment}
	ID, err := u.dbRepo.InsertComment(newComment)
	if err != nil {
		return nil, err
	}
	newComment.ID = ID

	return newComment, nil
}

// GetByOrg get all comments against a Github org
func (u *commentUsecase) GetByOrg(org string) ([]*domain.Comment, error) {
	return u.dbRepo.FetchCommentsByOrg(org)
}

// DeleteByOrg mark all comments against a Github org as deleted
func (u *commentUsecase) DeleteByOrg(org string) error {
	return u.dbRepo.DeleteCommentsByOrg(org)
}
