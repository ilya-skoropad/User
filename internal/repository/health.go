package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type HealthRepository interface {
	Ping() error
}

type health struct {
	connection *sql.DB
}

func (h *health) Ping() error {
	_, err := h.connection.Query("select 1")
	if err != nil {
		return err
	}

	return nil
}

func NewHealthRepository(connection *sql.DB) HealthRepository {
	return &health{
		connection: connection,
	}
}
