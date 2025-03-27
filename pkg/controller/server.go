package controller

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/repository"
)

type RestController struct {
	Repository repository.INovel
}

var restLock = &sync.Mutex{}

var RestControllerInstance *RestController

func GetRestControllerInstance() *RestController {
	if RestControllerInstance == nil {
		restLock.Lock()
		defer restLock.Unlock()

		if RestControllerInstance == nil {
			repo, err := repository.CreateNovelRepo()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil
			}

			RestControllerInstance = &RestController{
				Repository: repo,
			}
		} else {
			fmt.Println("Novel controller instance already created")
		}
	} else {
		fmt.Println("Novel controler instance already created")
	}

	return RestControllerInstance
}
