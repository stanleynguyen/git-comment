package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stanleynguyen/git-comment/comment-app/delivery/http/comment"
	"github.com/stanleynguyen/git-comment/comment-app/usecase"
)

// Handler http routes handler
type Handler struct {
	Router *httprouter.Router
}

// InitCommentsHandler hook up handler fn(s) with router
func (h *Handler) InitCommentsHandler(cu usecase.CommentUsecase) *Handler {
	commentHandler := &comment.Handler{CommentUsecase: cu}
	h.Router.GET("/orgs/:org/comments", commentHandler.GetCommentsByOrg)
	h.Router.POST("/orgs/:org/comments", commentHandler.CreateComment)
	h.Router.DELETE("/orgs/:org/comments", commentHandler.DeleteComment)
	return h
}
