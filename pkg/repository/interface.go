package repository

type INovel interface {
	CreateNovel() error
	DeleteNovel() error
	UpdateNovel() error
	GetNovel() error
	GetAllNovel() error
}
