package persistence

import (
	"database/sql"
	"fmt"

	"log"

	"github.com/stanleynguyen/git-comment/comment-app/domain"
	"github.com/stanleynguyen/git-comment/comment-app/repository"
)

type postgresPersistence struct {
	Conn *sql.DB
}

// NewPostgresRepo generate new Postgres client
func NewPostgresRepo(Conn *sql.DB) repository.Persistence {
	return &postgresPersistence{Conn}
}

// FetchCommentsByOrg query all comments against an org
func (p *postgresPersistence) FetchCommentsByOrg(org string) ([]*domain.Comment, error) {
	query := `SELECT id, org, comment FROM comments WHERE org = $1 AND NOT deleted`
	comments, err := p.fetchComments(query, org)
	if err != nil {
		logDBErr(err)
		return nil, domain.NewErrorInternalServer(
			fmt.Sprintf("Fail to fetch comments against %s", org),
		)
	}

	return comments, nil
}

// InsertComment insert a comment into comments table
func (p *postgresPersistence) InsertComment(c *domain.Comment) (int64, error) {
	query := `INSERT INTO comments(org, comment) VALUES ($1, $2) RETURNING id`
	var ID int64
	err := p.Conn.QueryRow(query, c.Org, c.Comment).Scan(&ID)
	if err != nil {
		logDBErr(err)
		return 0, domain.NewErrorInternalServer("Fail to create comment")
	}

	return ID, nil
}

// DeleteCommentsByOrg soft-delete all comments against an org
func (p *postgresPersistence) DeleteCommentsByOrg(org string) error {
	query := `UPDATE comments SET deleted = TRUE WHERE org = $1 AND NOT deleted`
	stmt, err := p.Conn.Prepare(query)
	ise := domain.NewErrorInternalServer(fmt.Sprintf("Fail to delete comments against %s", org))
	if err != nil {
		logDBErr(err)
		return ise
	}

	res, err := stmt.Exec(org)
	if err != nil {
		logDBErr(err)
		return ise
	}

	_, err = res.RowsAffected()
	if err != nil {
		logDBErr(err)
		return ise
	}

	return nil
}

func (p *postgresPersistence) fetchComments(query string, args ...interface{}) ([]*domain.Comment, error) {
	rows, err := p.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*domain.Comment, 0)
	for rows.Next() {
		c := new(domain.Comment)
		err = rows.Scan(&c.ID, &c.Org, &c.Comment)

		if err != nil {
			return nil, err
		}
		results = append(results, c)
	}

	return results, nil
}

func logDBErr(err error) {
	log.Printf("DB error: %s", err.Error())
}
