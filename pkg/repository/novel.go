package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/domain"
	"github.com/vitorsavian/tracker/pkg/infra/database"
)

type NovelRepositoryDB struct {
	Driver database.IConnection
}

func CreateNovelRepo() (*NovelRepositoryDB, error) {
	driver := &database.PSQL{}
	if err := driver.CreateConnection(); err != nil {
		return nil, err
	}

	novelRepo := &NovelRepositoryDB{}

	novelRepo.Driver = driver
	return novelRepo, nil
}

func (n *NovelRepositoryDB) CreateNovel(novel *domain.Novel) error {
	request := adapter.CreateNovelDatabaseRequestAdapter{
		Id:       novel.Id,
		Name:     novel.Name,
		Finished: novel.Finished,
		Page:     novel.Page,
	}

	err := n.Driver.CreateNovel(&request)
	if err != nil {
		return err
	}

	return nil
}

func (n *NovelRepositoryDB) DeleteNovel(id string) error {
	err := n.Driver.DeleteNovel(id)
	if err != nil {
		return err
	}

	return nil
}

func (n *NovelRepositoryDB) UpdateNovel(novel *domain.Novel) error {
	request := adapter.UpdateNovelDatabaseRequestAdapter{
		Id:       novel.Id,
		Name:     novel.Name,
		Finished: novel.Finished,
		Page:     novel.Page,
	}

	err := n.Driver.UpdateNovel(&request)
	if err != nil {
		return err
	}

	return nil
}

func (n *NovelRepositoryDB) GetNovel(id string) (*domain.Novel, error) {
	response, err := n.Driver.GetNovel(id)
	if err != nil {
		return nil, err
	}

	novel := &domain.Novel{
		Id:       response.Id,
		Name:     response.Name,
		Finished: response.Finished,
		Page:     response.Page,
	}

	return novel, nil
}

func (n *NovelRepositoryDB) GetAllNovel() ([]domain.Novel, error) {
	response, err := n.Driver.GetAllNovel()
	if err != nil {
		logrus.Errorf("unable to get novels: %v\n", err)
		return nil, err
	}
	fmt.Println(response)

	var novels []domain.Novel
	for _, v := range response {
		fmt.Println(v)
		novel := domain.Novel{
			Id:       v.Id,
			Name:     v.Name,
			Finished: v.Finished,
			Page:     v.Page,
		}

		novels = append(novels, novel)
	}

	return novels, nil
}
