package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vitorsavian/tracker/pkg/adapter"
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
				logrus.Errorf("Unable to novel instance: %v\n", err)
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "600")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *Controller) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/novel", c.NovelEndpoint)
	handler := corsMiddleware(mux)

	server := http.Server{
		Addr:                         os.Getenv("SERVER_ADDR"),
		Handler:                      handler,
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
		logrus.Infof("Starting server on %s", server.Addr)
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

func setErrorResponse(w http.ResponseWriter, status int, err error) {
	if errResp := json.NewEncoder(w).Encode(adapter.Response{
		Status:  status,
		Message: err.Error(),
	}); errResp != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("Error encoding error response: %v", errResp)
	}
}

func (c *Controller) NovelEndpoint(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Method received: %s", r.Method)
	logrus.Infof("URL: %s", r.URL.String())
	logrus.Infof("Headers: %+v", r.Header)

	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id == "" {
			c.GetAllNovel(w, r)
			return
		}
		c.GetNovel(w, r)
		return
	case http.MethodPost:
		c.CreateNovel(w, r)
		return
	case http.MethodPut:
		c.UpdateNovel(w, r)
		return
	case http.MethodDelete:
		c.DeleteNovel(w, r)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Controller) GetNovel(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("GetNovel called")
	w.Header().Set("Content-Type", "application/json")

	novel, status, err := c.Novel.GetNovel(r.URL.Query().Get("id"))
	if err != nil {
		logrus.Errorf("Error getting novel: %v", err)
		setErrorResponse(w, status, err)
		return
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(novel); err != nil {
		logrus.Errorf("Error encoding novel: %v", err)
		setErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (c *Controller) GetAllNovel(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("GetAllNovel called")
	w.Header().Set("Content-Type", "application/json")

	novels, status, err := c.Novel.GetAllNovel()
	if err != nil {
		logrus.Errorf("Error getting all novels: %v", err)
		setErrorResponse(w, status, err)
		return
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(novels); err != nil {
		logrus.Errorf("Error encoding novels: %v", err)
		setErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (c *Controller) CreateNovel(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("CreateNovel called")
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body: %v", err)
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()
	data := &adapter.CreateNovelAdapter{}

	if err := json.Unmarshal(body, data); err != nil {
		logrus.Errorf("Error unmarshalling request body: %v", err)
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	novel, status, err := c.Novel.CreateNovel(data)
	if err != nil {
		logrus.Errorf("Error creating novel: %v", err)
		setErrorResponse(w, status, err)
		return
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(novel); err != nil {
		logrus.Errorf("Error encoding novel: %v", err)
		setErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (c *Controller) UpdateNovel(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("UpdateNovel called")
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		setErrorResponse(w, http.StatusBadRequest, errors.New(
			"ID is required for updating a novel",
		))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body: %v", err)
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()
	data := &adapter.UpdateNovelAdapter{}
	if err := json.Unmarshal(body, data); err != nil {
		logrus.Errorf("Error unmarshalling request body: %v", err)
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	data.Id = id

	_, status, err := c.Novel.UpdateNovel(data)
	if err != nil {
		logrus.Errorf("Error updating novel: %v", err)
		setErrorResponse(w, status, err)
		return
	}

	w.WriteHeader(status)
}

func (c *Controller) DeleteNovel(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("DeleteNovel called")
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		setErrorResponse(w, http.StatusBadRequest, errors.New(
			"ID is required for deleting a novel",
		))
		return
	}

	status, err := c.Novel.DeleteNovel(id)
	if err != nil {
		logrus.Errorf("Error deleting novel: %v", err)
		setErrorResponse(w, status, err)
		return
	}

	w.WriteHeader(status)
}
