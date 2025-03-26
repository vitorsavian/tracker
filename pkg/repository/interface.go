package repository

import "github.com/vitorsavian/tracker/pkg/domain"

type INovel interface {
	CreateNovel(novel *domain.Novel) error
	DeleteNovel() error
	UpdateNovel() error
	GetNovel() error
	GetAllNovel() error
}
