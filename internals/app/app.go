package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/c0rnoTa/goagi/internals/app/models/repository/mysql"
	"github.com/c0rnoTa/goagi/internals/app/processor"
	"github.com/c0rnoTa/goagi/internals/cfg"
	"log"
	"net"
)

// will be filled at build phase
var gitHash, buildTime string

const appName = "AGI"

type Server struct {
	debug   bool
	config  *cfg.Configuration
	storage *mysql.Storage
}

func NewServer(cfg *cfg.Configuration) *Server {
	return &Server{
		debug:  cfg.App.Debug,
		config: cfg,
	}
}

func (s *Server) Serve(ctx context.Context) error {
	version := fmt.Sprintf("(%s build at %s)", gitHash, buildTime)
	if gitHash == "" {
		version = "(develop)"
	}
	log.Println(appName, version, "server starting")

	// Подключаем базу данных
	databaseAddress := fmt.Sprintf("tcp(%s)", net.JoinHostPort(s.config.Database.Host, fmt.Sprintf("%d", s.config.Database.Port)))
	if s.debug {
		log.Printf("Connecting to database %s@%s:%d/%s", s.config.Database.Login, s.config.Database.Host, s.config.Database.Port, s.config.Database.Name)
	}
	if s.storage = mysql.NewStorage(ctx, s.config.Database.Login, s.config.Database.Password, databaseAddress, s.config.Database.Name); s.storage == nil {
		return errors.New("failed to create new storage")
	}

	// TODO: обработка здесь
	proc := processor.NewProcessor(s.storage, nil)
	proc.BlacklistCheck(ctx)
	log.Println(appName, "server started")
	return nil
}

func (s *Server) Shutdown() {
	log.Println(appName, "server is shutting down")
	if s.debug {
		log.Printf("Disconnecting from database")
	}
	if err := s.storage.DB.Close(); err != nil {
		log.Println("Unable to disconnect from database:", err)
	}
	log.Println(appName, "stopped!")
}
