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

	err = c.Repository.CreateNovel(novel)
	if err != nil {
		return err
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("Novel id: %s\n", novel.Id)
	fmt.Printf("Novel name: %s\n", novel.Name)
	fmt.Printf("Novel page: %d\n", novel.Page)
	fmt.Printf("Novel finished: %t\n", novel.Finished)
	fmt.Println("--------------------------------------")

	return nil
}

func (c *NovelController) CliDelete(id string) error {
	err := c.Repository.DeleteNovel(id)
	if err != nil {
		return err
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("Novel deleted: %s\n", id)
	fmt.Println("--------------------------------------")

	return nil
}

func (c *NovelController) CliUpdate(adapter *adapter.UpdateNovelAdapter) error {
	novel, err := domain.UpdateNovel(adapter)
	if err != nil {
		return err
	}

	err = c.Repository.UpdateNovel(novel)
	if err != nil {
		return err
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("Novel updated: %s\n", adapter.Id)
	fmt.Println("--------------------------------------")
	fmt.Printf("With this values:\n")
	fmt.Printf("Novel name: %s\n", novel.Name)
	fmt.Printf("Novel page: %d\n", novel.Page)
	fmt.Printf("Novel finished: %t\n", novel.Finished)
	fmt.Println("--------------------------------------")

	return nil
}

func (c *NovelController) CliGet(id string) error {
	novel, err := c.Repository.GetNovel(id)
	if err != nil {
		return err
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("Novel id: %s\n", novel.Id)
	fmt.Printf("Novel name: %s\n", novel.Name)
	fmt.Printf("Novel page: %d\n", novel.Page)
	fmt.Printf("Novel finished: %t\n", novel.Finished)
	fmt.Println("--------------------------------------")

	return nil
}

func (c *NovelController) CliGetAll() error {
	novels, err := c.Repository.GetAllNovel()
	if err != nil {
		return err
	}

	for _, v := range novels {
		fmt.Println("--------------------------------------")
		fmt.Printf("Novel id: %s\n", v.Id)
		fmt.Printf("Novel name: %s\n", v.Name)
		fmt.Printf("Novel page: %d\n", v.Page)
		fmt.Printf("Novel finished: %t\n", v.Finished)
	}
	fmt.Println("--------------------------------------")

	return nil
}
