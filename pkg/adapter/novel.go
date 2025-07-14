package adapter

import "time"

type CreateNovelAdapter struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	Finished bool   `json:"finished"`
}

type UpdateNovelAdapter struct {
	Id       string
	Name     string `json:"name"`
	Page     int    `json:"page"`
	Finished bool   `json:"finished"`
}

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
	UpdatedAt *time.Time
}
