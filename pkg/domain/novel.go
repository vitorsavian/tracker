package domain

import (
	"errors"

	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/utils"
)

type Novel struct {
	Id       string
	Name     string
	Page     int
	Finished bool
}

func NewNovel(novel *adapter.CreateNovelAdapter) (*Novel, error) {
	if novel.Name == "" {
		return nil, errors.New("novel with blank name")
	}

	if novel.Page < 0 {
		return nil, errors.New("novel with pages below 0")
	}

	return &Novel{
		Id:       utils.GenerateUUID(),
		Name:     novel.Name,
		Page:     novel.Page,
		Finished: novel.Finished,
	}, nil
}

func UpdateNovel(novel *adapter.UpdateNovelAdapter) (*Novel, error) {
	if novel.Name == "" {
		return nil, errors.New("novel with blank name")
	}

	if novel.Page < 0 {
		return nil, errors.New("novel with pages below 0")
	}

	return &Novel{
		Id:       novel.Id,
		Name:     novel.Name,
		Page:     novel.Page,
		Finished: novel.Finished,
	}, nil
}

func GetAllPages(novel []adapter.GetNovelDatabaseResponseAdapter) int {
	pages := 0
	for _, v := range novel {
		pages += v.Page
	}

	return pages
}
