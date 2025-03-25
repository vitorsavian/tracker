package domain

import "github.com/vitorsavian/tracker/pkg/adapter"

type Novel struct {
	Id       string
	Name     string
	Page     int
	Finished bool
}

func CreateNovel(novel *adapter.CreateNovelAdapter) error {
	return nil
}
