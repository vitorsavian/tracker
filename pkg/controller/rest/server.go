package rest

import (
	"fmt"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/domain"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/repository"
)

type Controller struct {
	Repository repository.INovel
}

var restLock = &sync.Mutex{}

var ControllerInstance *Controller

func GetControllerInstance() *Controller {
	if ControllerInstance == nil {
		restLock.Lock()
		defer restLock.Unlock()

		if ControllerInstance == nil {
			repo, err := repository.CreateNovelRepo()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil
			}

			ControllerInstance = &Controller{
				Repository: repo,
			}
		} else {
			fmt.Println("Novel controller instance already created")
		}
	} else {
		fmt.Println("Novel controler instance already created")
	}

	return ControllerInstance
}

func (c *Controller) CreateNovel(requestAdapter *adapter.CreateNovelAdapter) (*domain.Novel, int, error) {
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

func (c *Controller) DeleteNovel(id string) (int, error) {
	err := c.Repository.DeleteNovel(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (c *Controller) UpdateNovel(adapter *adapter.UpdateNovelAdapter) (*domain.Novel, int, error) {
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

func (c *Controller) GetNovel(id string) (*domain.Novel, int, error) {
	novel, err := c.Repository.GetNovel(id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return novel, http.StatusOK, nil
}

func (c *Controller) GetAllNovel() (*[]domain.Novel, int, error) {
	novels, err := c.Repository.GetAllNovel()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &novels, http.StatusOK, nil
}
