package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/domain"
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

var createNovel = `
INSERT INTO novel(id, name, page, finished) VALUES($1, $2, $3, $4)
`

func (p *PSQL) CreateNovel(novel *domain.Novel) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin transction: %v\n", err)
		return err
	}

	if _, err = tx.Exec(ctx, createNovel, novel.Id,
		novel.Name, novel.Page, novel.Finished); err != nil {
		logrus.Errorf("Unable to exec query to create novel: %v\n", err)

		if err = tx.Rollback(ctx); err != nil {
			logrus.Errorf("Unable to rollback create novel transaction: %v\n", err)
			return err
		}

		return err
	}

	tx.Commit(ctx)
	return nil
}
