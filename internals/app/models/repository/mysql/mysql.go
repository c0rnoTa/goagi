package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Storage struct {
	DB *sqlx.DB
}

const (
	driverName            = "mysql"
	connectionTimeout     = 5 * time.Second
	connectionMaxLifetime = 3 * time.Minute
	connectionsMax        = 10
)

func NewStorage(ctx context.Context, username, password, address, databaseName string) *Storage {

	// Create a connection string.
	mysqlDsn := fmt.Sprintf("%s:%s@%s/%s", username, password, address, databaseName)

	// Connect to server..
	db, err := sqlx.Open(driverName, mysqlDsn)
	if err != nil {
		log.Printf("Failed to verify database connection string: %v\n", err)
		return nil
	}

	db.SetConnMaxLifetime(connectionMaxLifetime)
	db.SetMaxOpenConns(connectionsMax)
	db.SetMaxIdleConns(connectionsMax)

	storage := Storage{DB: db}
	if err = storage.checkConnection(ctx); err != nil {
		storage.DB.Close()
		log.Printf("Unable to connect to database: %v\n", err)
		return nil
	}

	return &storage
}

// checkConnection checks if connection still successful
func (s *Storage) checkConnection(ctx context.Context) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()
	if err := s.DB.PingContext(ctxTimeout); err != nil {
		return fmt.Errorf("connection check failed: %w", err)
	}
	return nil
}
