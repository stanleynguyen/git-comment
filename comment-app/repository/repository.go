package repository

import "github.com/stanleynguyen/git-comment/comment-app/domain"

// Persistence abstraction against database lib
type Persistence interface {
	FetchCommentsByOrg(org string) ([]*domain.Comment, error)
	InsertComment(comment *domain.Comment) (int64, error)
	DeleteCommentsByOrg(org string) error
}

// GithubCli github API client
type GithubCli interface {
	OrgExists(org string) (bool, error)
}
