package internal

import (
	"fmt"
	"github.com/0x9p/coding_task_1/internal/domain/port"
	"github.com/0x9p/coding_task_1/internal/nap"
	portRoute "github.com/0x9p/coding_task_1/internal/route/v1/port"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App interface {
	Init()
	Shutdown()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type app struct {
	config *Config
	router *mux.Router
	db     nap.SqlDb
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *app) Init() {
	db, err := nap.Open("postgres", []string{
		fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s sslmode=disable",
			a.config.MasterDb.Host,
			a.config.MasterDb.Name,
			a.config.MasterDb.User,
			a.config.MasterDb.Password,
		),
	})

	if err != nil {
		log.Fatal(err)
	}

	a.db = db

	if err := db.Ping(); err != nil {
		log.Fatalf("Some physical database is unreachable: %s", err)
	}

	portRepo := port.NewRepo(db)
	portService := port.NewService(portRepo)
	portHandler := portRoute.NewHandler(portService)

	portRoute.Route(a.router, portHandler)
}

func (a *app) Shutdown() {
	if err := a.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func NewApp(config *Config) App {
	return &app{
		config: config,
		router: mux.NewRouter()}
}
