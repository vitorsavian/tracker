package controller

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/adapter"
	"github.com/vitorsavian/tracker/pkg/domain"
	"github.com/vitorsavian/tracker/pkg/repository"
)

type NovelController struct {
	Repository repository.INovel
}

var novelLock = &sync.Mutex{}

var NovelControllerInstance *NovelController

func GetNovelControllerInstance() *NovelController {
	if NovelControllerInstance == nil {
		novelLock.Lock()
		defer novelLock.Unlock()

		if NovelControllerInstance == nil {
			repo, err := repository.CreateNovelRepo()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil
			}

			NovelControllerInstance = &NovelController{
				Repository: repo,
			}
		} else {
			fmt.Println("Novel controller instance already created")
		}
	} else {
		fmt.Println("Novel controler instance already created")
	}

	return NovelControllerInstance
}

func (c *NovelController) CliCreate(adapter *adapter.CreateNovelAdapter) error {
	novel, err := domain.NewNovel(adapter)
	if err != nil {
		return err
	}

	fmt.Println(novel)

	return nil
}
