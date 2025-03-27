package database

import (
	"github.com/vitorsavian/tracker/pkg/adapter"
)

type IConnection interface {
	CreateConnection() error

	CreateNovel(novel *adapter.CreateNovelDatabaseRequestAdapter) error
	DeleteNovel(id string) error
	UpdateNovel(novel *adapter.UpdateNovelDatabaseRequestAdapter) error
	GetNovel(id string) (*adapter.GetNovelDatabaseResponseAdapter, error)
	GetAllNovel() ([]adapter.GetNovelDatabaseResponseAdapter, error)
}
