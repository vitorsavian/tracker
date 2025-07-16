package rest

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/vitorsavian/tracker/pkg/usecase"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	Novel *usecase.Novel
}

var restLock = &sync.Mutex{}

var ControllerInstance *Controller

func GetControllerInstance() *Controller {
	if ControllerInstance == nil {
		restLock.Lock()
		defer restLock.Unlock()

		if ControllerInstance == nil {
			novelInstance, err := usecase.GetNovelInstance()
			if err != nil {
				logrus.Errorf("Unable to create repository: %v\n", err)
				return nil
			}

			ControllerInstance = &Controller{
				Novel: novelInstance,
			}
		} else {
			logrus.Infoln("Rest controller instance already created")
		}
	} else {
		logrus.Infoln("Rest controller instance already created")
	}

	return ControllerInstance
}

func (c *Controller) Start() {

	mux := http.NewServeMux()
	mux.HandleFunc("/novel", c.NovelEndpoint)

	server := http.Server{
		Addr:                         os.Getenv("SERVER_ADDR"),
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  60 * time.Second,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     log.New(logrus.StandardLogger().Out, "server: ", log.LstdFlags),
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Server failed: %v", err)
		}
	}()

	// Listen for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Infoln("Server exited gracefully")
}

func (c *Controller) NovelEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.GetNovel(w, r)
	case http.MethodPost:
		c.CreateNovel(w, r)
	case http.MethodPut:
		c.UpdateNovel(w, r)
	case http.MethodDelete:
		c.DeleteNovel(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Controller) GetNovel(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to get a novel
}

func (c *Controller) CreateNovel(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to create a novel
}

func (c *Controller) UpdateNovel(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update a novel
}

func (c *Controller) DeleteNovel(w http.ResponseWriter, r *http.Request) {

}
