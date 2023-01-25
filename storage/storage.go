package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/uzbekman2005/real_time_chatting/storage/postgres"
	"github.com/uzbekman2005/real_time_chatting/storage/repo"
)

type IStorage interface {
	Chat() repo.ChatStorageI
}

type StoragePg struct {
	Db       *sqlx.DB
	chatRepo repo.ChatStorageI
}

// NewStoragePg
func NewStoragePg(db *sqlx.DB) *StoragePg {
	return &StoragePg{
		Db:       db,
		chatRepo: postgres.NewChatRepo(db),
	}
}

func (s StoragePg) Chat() repo.ChatStorageI {
	return s.chatRepo
}
