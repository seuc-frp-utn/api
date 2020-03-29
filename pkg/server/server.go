package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Config represents a set of parameters to initialize a Server
type Config struct {
	Host string
	Port uint
	Router *gin.Engine
}

// Server wraps a http server
type Server http.Server

// New initializes a new Server by the given configuration.
func New(config Config) *Server {
	return &Server{
		Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:           config.Router,
		ReadTimeout:       time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
	}
}

