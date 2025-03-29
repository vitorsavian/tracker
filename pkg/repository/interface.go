package repository

import "github.com/vitorsavian/tracker/pkg/domain"

type INovel interface {
	CreateNovel(novel *domain.Novel) error
	DeleteNovel(id string) error
	UpdateNovel(novel *domain.Novel) error
	GetNovel(id string) (*domain.Novel, error)
	GetAllNovel() ([]domain.Novel, error)
}

type IHealth interface {
	CreateCaloriesLog() error
	DeleteCaloriesLog() error
	UpdateCaloriesLog() error
	GetCaloriesLog() error
}
