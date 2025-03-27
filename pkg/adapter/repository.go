package adapter

import "time"

type CreateNovelDatabaseRequestAdapter struct {
	Id       string
	Name     string
	Page     int
	Finished bool
}

type UpdateNovelDatabaseRequestAdapter struct {
	Id       string
	Name     string
	Page     int
	Finished bool
}

type GetNovelDatabaseResponseAdapter struct {
	Id        string
	Name      string
	Page      int
	Finished  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
