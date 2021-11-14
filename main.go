package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	plogger "gitlab.com/protocole/clearkey/internal/core/ports/logger"
	"gitlab.com/protocole/clearkey/internal/core/ports/repositories"
	"gitlab.com/protocole/clearkey/internal/core/services"
	"gitlab.com/protocole/clearkey/internal/handlers"
	zlogger "gitlab.com/protocole/clearkey/internal/loggers"
	"gitlab.com/protocole/clearkey/internal/repositories/memory"
	"gitlab.com/protocole/clearkey/internal/repositories/postgresql"
	"gitlab.com/protocole/clearkey/pkg/apperrors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := loadEnvironment()
	if err != nil {
		log.Fatal(err)
		return
	}
	// only choice available for now is zlogger
	plogger.SetLogger(zlogger.NewZLogger(viper.GetString("EnvType")))

	service := services.NewService(chooseRepository())
	handler := handlers.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: viper.GetStringSlice("Domains"),
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/license", handler.GetKey)
	r.Post("/license/register", handler.PostKey)

	errs := make(chan error, 2)
	go func() {
		plogger.Log.Infof("Application is now listening on %s:%s", viper.GetString("Ip"), viper.GetString("Port"))
		ip := fmt.Sprintf("%s:%s", viper.GetString("Ip"), viper.GetString("Port"))
		errs <- http.ListenAndServe(ip, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	plogger.Log.Errorf(fmt.Sprintf("\nTerminated %s\n", <-errs))
}

/*
  This function use the native logger as we load the logger after this method
*/
func loadEnvironment() error {
	// This corresponds to key ID & env var name
	envKeys := map[string]string{
		"EnvType":       "ENV",
		"Port":          "PORT",
		"Ip":            "IP",
		"Domains":       "ALLOWED_DOMAINS",
		"Repository":    "REPOSITORY_TYPE",
		"Psql_pass":     "PSQL_PASSWORD",
		"Psql_user":     "PSQL_USER",
		"Psql_addr":     "PSQL_ADDR",
		"Psql_db":       "PSQL_DB",
		"Psql_insecure": "PSQL_INSECURE",
	}

	for keyId, keyValue := range envKeys {
		if err := viper.BindEnv(keyId, keyValue); err != nil {
			log.Fatalf("Failed to boot up environment variables, reason: %s", err.Error())
			return apperrors.EnvVarLoadFailed
		}
	}

	viper.SetDefault("EnvType", "development")
	viper.SetDefault("Port", 8080)
	viper.SetDefault("Ip", "0.0.0.0")
	viper.SetDefault("ALLOWED_DOMAINS", []string{"http://localhost:*", "http://127.0.0.1:*"})
	viper.SetDefault("Repository", "RAM")
	viper.SetDefault("Psql_pass", "postgres")
	viper.SetDefault("Psql_user", "postgres")
	viper.SetDefault("Psql_addr", "127.0.0.1:5433")
	viper.SetDefault("Psql_db", "postgres")
	viper.SetDefault("Psql_insecure", true)

	return nil
}

func chooseRepository() repositories.KeyStorageRepository {
	repositoryType := viper.GetString("Repository")
	plogger.Log.Infof("Selected repository is %s", repositoryType)

	switch repositoryType {
	case "RAM":
		return memory.NewMemoryRepository()
	case "PSQL":
		repository, err := postgresql.NewPostgreSQLRepository(
			viper.GetString("Psql_user"),
			viper.GetString("Psql_pass"),
			viper.GetString("Psql_db"),
			viper.GetString("Psql_addr"),
			viper.GetBool("Psql_insecure"),
		)
		if err != nil {
			plogger.Log.Fatalf(err.Error())
		}

		return repository
	}

	plogger.Log.Fatalf("Invalid repository %s, please choose between the available ones", repositoryType)
	return nil
}
