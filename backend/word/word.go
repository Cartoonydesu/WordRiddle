package word

import "database/sql"

type Handler struct {
	Db *sql.DB
}

type Word struct {
	Word        string `json:"word" binding:"requried"`
	Description string `json:"description"`
}
