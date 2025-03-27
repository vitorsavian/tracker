package repository

import (
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
	err := n.Driver.CreateNovel(novel)
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

func (n *NovelRepositoryDB) UpdateNovel() error {
	return nil
}

func (n *NovelRepositoryDB) GetNovel() error {
	return nil
}

func (n *NovelRepositoryDB) GetAllNovel() error {
	return nil
}
