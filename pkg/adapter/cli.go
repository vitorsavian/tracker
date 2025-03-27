package adapter

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
