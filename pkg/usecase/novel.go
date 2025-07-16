package usecase

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/domain"
	"github.com/vitorsavian/tracker/pkg/repository"
	"net/http"
	"sync"
)

type Novel struct {
	Repository repository.INovel
}

var novelLock = &sync.Mutex{}

var novelInstance *Novel

func GetNovelInstance() (*Novel, error) {
	if novelInstance == nil {
		novelLock.Lock()
		defer novelLock.Unlock()

		if novelInstance == nil {
			repo, err := repository.CreateNovelRepo()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil, err
			}

			novelInstance = &Novel{
				Repository: repo,
			}
		} else {
			fmt.Println("Novel instance already created")
		}
	} else {
		fmt.Println("Novel instance already created")
	}

	return novelInstance, nil
}

func (c *Novel) CreateNovel(requestAdapter *adapter.CreateNovelAdapter) (*domain.Novel, int, error) {
	novel, err := domain.NewNovel(requestAdapter)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = c.Repository.CreateNovel(novel)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return novel, http.StatusCreated, nil
}

func (c *Novel) DeleteNovel(id string) (int, error) {
	err := c.Repository.DeleteNovel(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (c *Novel) UpdateNovel(adapter *adapter.UpdateNovelAdapter) (*domain.Novel, int, error) {
	novel, err := domain.UpdateNovel(adapter)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = c.Repository.UpdateNovel(novel)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return novel, http.StatusOK, nil
}

func (c *Novel) GetNovel(id string) (*domain.Novel, int, error) {
	novel, err := c.Repository.GetNovel(id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return novel, http.StatusOK, nil
}

func (c *Novel) GetAllNovel() (*[]domain.Novel, int, error) {
	novels, err := c.Repository.GetAllNovel()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &novels, http.StatusOK, nil
}
