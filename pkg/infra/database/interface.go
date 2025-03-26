package database

import "github.com/vitorsavian/tracker/pkg/domain"

type IConnection interface {
	CreateConnection() error
	CreateNovel(novel *domain.Novel) error
}
