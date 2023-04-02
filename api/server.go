package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hoangtk0100/go-healthy/db/sqlc"
	"github.com/hoangtk0100/go-healthy/util"
)

var currentUserName = "user@demo.com"

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/meals", server.createMeal)
	router.GET("/meals", server.listMeals)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
