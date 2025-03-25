package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PSQL struct {
	conn *pgxpool.Pool
}

func (p *PSQL) CreateConnection() error {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Errorf("Unable to create connection pool: %v\n", err)
		return err
	}

	p.conn = conn

	return nil
}
