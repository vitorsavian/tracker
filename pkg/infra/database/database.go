package database

type INovel interface {
	CreateNovel() error
	DeleteNovel() error
	UpdateNovel() error
	GetNovel() error
	GetAllNovel() error
}

type IHealth interface {
	CreateHealth() error
}
