package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	logger2 "gitlab.com/protocole/clearkey/internal/core/ports/logger"
	services2 "gitlab.com/protocole/clearkey/internal/core/services"
	handlers2 "gitlab.com/protocole/clearkey/internal/handlers"
	loggers2 "gitlab.com/protocole/clearkey/internal/loggers"
	repositories2 "gitlab.com/protocole/clearkey/internal/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger2.SetLogger(loggers2.NewZLogger())
	repo := repositories2.NewMemoryRepository()
	service := services2.NewService(repo)
	handler := handlers2.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Post("/license", handler.GetKey)
	r.Post("/license/register", handler.PostKey)

	errs := make(chan error, 2)
	go func() {
		logger2.Log.Infof("Listening on prt :8080")
		errs <- http.ListenAndServe(":8080", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger2.Log.Errorf(fmt.Sprintf("\nTerminated %s\n", <-errs))
}
