package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type ChatRepo struct {
	Db *sqlx.DB
}

func NewChatRepo(db *sqlx.DB) *ChatRepo {
	return &ChatRepo{Db: db}
}
