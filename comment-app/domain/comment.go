package domain

// Comment domain object for comment saved in PostgreSQL
type Comment struct {
	ID        int64     `json:"id,omitempty"`
	Org       string    `json:"org,omitempty"`
	Comment   string    `json:"comment" schema:"comment"`
}
