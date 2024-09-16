package httpServer

import (
	"go-backend/src/config"
	"go-backend/src/logger"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/jmoiron/sqlx"
)

type HttpServer struct {
	Url    string
	DB     *sqlx.DB
	Router *chi.Mux
}

func InitServer(cfg config.Config, Db *sqlx.DB) *HttpServer {
	address := cfg.PublicHost + ":" + cfg.Port
	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(logger.GetLogger()))

	return &HttpServer{
		Url:    address,
		DB:     Db,
		Router: router,
	}
}

func (server *HttpServer) Run() {
	err := http.ListenAndServe(server.Url, server.Router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Print(server.Url)
}
