package adapter

type CreateNovelAdapter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Page     int    `json:"page"`
	Finished bool   `json:"finished"`
}

type UpdateNovelAdapter struct {
	Id       string
	Name     string
	Page     int
	Finished bool
}
