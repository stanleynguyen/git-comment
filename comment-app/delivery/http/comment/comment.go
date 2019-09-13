package comment

import (
	"github.com/stanleynguyen/git-comment/comment-app/delivery/http/helpers"
	"github.com/stanleynguyen/git-comment/comment-app/domain"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stanleynguyen/git-comment/comment-app/usecase"
)

// Handler comment routes handler
type Handler struct {
	CommentUsecase usecase.CommentUsecase
}

// CreateComment POST /orgs/:org/comments
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	org := ps.ByName("org")
	c := new(domain.Comment)
	if err := helpers.ReadRequestBody(r, c); err != nil {
		helpers.RenderErr(w, err)
		return
	}
	if len(c.Comment) > 255 || len(c.Comment) == 0 {
		helpers.RenderErr(w, domain.NewErrorUnprocessableEntity("comment's length most be within [0-255] characters"))
		return
	}
	comment, err := h.CommentUsecase.Create(org, c.Comment)
	if err != nil {
		helpers.RenderErr(w, err)
		return
	}

	helpers.RenderJSON(w, comment)
}

// GetCommentsByOrg GET /orgs/:org/comments
func (h *Handler) GetCommentsByOrg(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	org := ps.ByName("org")
	comments, err := h.CommentUsecase.GetByOrg(org)
	if err != nil {
		helpers.RenderErr(w, err)
		return
	}

	helpers.RenderJSON(w, map[string][]*domain.Comment{"comments": comments})
}

// DeleteComment DELETE /orgs/:org/comments
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	org := ps.ByName("org")
	err := h.CommentUsecase.DeleteByOrg(org)
	if err != nil {
		helpers.RenderErr(w, err)
		return
	}

	helpers.RenderJSON(w, map[string]string{})
}
