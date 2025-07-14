package usecase

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/repository"
	"sync"
)

type Novel struct {
	Repository repository.INovel
}

var novelLock = &sync.Mutex{}

var novelInstance *Novel

func GetNovelInstance() *Novel {
	if novelInstance == nil {
		novelLock.Lock()
		defer novelLock.Unlock()

		if novelInstance == nil {
			repo, err := repository.CreateNovelRepo()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil
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

	return novelInstance
}
