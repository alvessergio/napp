package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/alvessergio/pan-integrations/application/repositories"
	"github.com/alvessergio/pan-integrations/application/services"
	"github.com/alvessergio/pan-integrations/framework/database"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort        = "8080"
	sideCarWaitingTime = 10 * time.Second
)

var db database.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		log.Fatalf("Error parsing boolean env var")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Error parsing boolean env var")
	}

	db.AutoMigrateDb = autoMigrateDb
	db.Debug = debug
	db.DsnTest = os.Getenv("DSN_TEST")
	db.Dsn = os.Getenv("DSN")
	db.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	db.DbType = os.Getenv("DB_TYPE")
	db.Env = os.Getenv("ENV")
}

func panicRecover() {
	if r := recover(); r != nil {
		log.WithFields(log.Fields{
			"event": "service_panicked",
		}).Errorf("error: %v", r)
	}
	log.Info("napp api server stopped")
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer panicRecover()

	go func() {
		l := log.WithFields(log.Fields{
			"port": defaultPort,
		})

		l.Info("starting napp api")

		dbConnection, err := db.Connect()

		if err != nil {
			log.Fatalf("error connecting to DB %v", err.Error())
		}

		productRepository := repositories.NewProductRepository(dbConnection)
		productHistoryRepository := repositories.NewProductHistoryRepository(dbConnection)

		svc := services.NewService(productRepository, productHistoryRepository)

		router := svc.NewServer()

		err = http.ListenAndServe(":"+defaultPort, router)
		if err != nil {
			log.Errorf("error while starting server. %v", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router))
			done <- syscall.SIGTERM
		}
	}()

	<-done
}
