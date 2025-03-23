package novel

import (
	"fmt"
	"sync"
)

type NovelController struct {
}

var novelLock = &sync.Mutex{}

var NovelControllerInstance *NovelController

func getNovelControllerInstance() *NovelController {
	if NovelControllerInstance == nil {
		novelLock.Lock()
		defer novelLock.Unlock()

		if NovelControllerInstance == nil {
			NovelControllerInstance = &NovelController{}
		} else {
			fmt.Println("Novel controller instance already created")
		}
	} else {
		fmt.Println("Novel controler instance already created")
	}

	return NovelControllerInstance
}
