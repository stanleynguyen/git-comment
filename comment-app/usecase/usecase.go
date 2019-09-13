package usecase

import "github.com/stanleynguyen/git-comment/comment-app/domain"

// CommentUsecase usecase contract for handling comment endpoints operations
type CommentUsecase interface {
	Create(org, comment string) (*domain.Comment, error)
	GetByOrg(org string) ([]*domain.Comment, error)
	DeleteByOrg(org string) error
}
