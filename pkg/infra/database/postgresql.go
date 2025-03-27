package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/domain"
	"github.com/vitorsavian/tracker/pkg/utils"
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
INSERT INTO novel(id, name, page, finished, created_at) VALUES($1, $2, $3, $4, $5)
`

func (p *PSQL) CreateNovel(novel *domain.Novel) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin transction: %v\n", err)
		return err
	}

	if _, err = tx.Exec(ctx, createNovel, novel.Id,
		novel.Name, novel.Page, novel.Finished, utils.UTCTime()); err != nil {
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

var updateNovel = `
UPDATE novel SET name = $2, page = $3, finished = $4, updated_at = $5 WHERE novel.id = $1;
`

func (p *PSQL) UpdateNovel(novel *domain.Novel) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin transction: %v\n", err)
		return err
	}

	if _, err = tx.Exec(ctx, updateNovel,
		novel.Id, novel.Name,
		novel.Page, novel.Finished, utils.UTCTime()); err != nil {
		logrus.Errorf("Unable to exec query to delete novel: %v\n", err)

		if err = tx.Rollback(ctx); err != nil {
			logrus.Errorf("Unable to rollback delete novel transaction: %v\n", err)
			return err
		}

		return err
	}

	tx.Commit(ctx)
	return nil
}

var deleteNovel = `
DELETE FROM novel WHERE novel.id = $1;
`

func (p *PSQL) DeleteNovel(id string) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin transction: %v\n", err)
		return err
	}

	if _, err = tx.Exec(ctx, deleteNovel, id); err != nil {
		logrus.Errorf("Unable to exec query to delete novel: %v\n", err)

		if err = tx.Rollback(ctx); err != nil {
			logrus.Errorf("Unable to rollback delete novel transaction: %v\n", err)
			return err
		}

		return err
	}

	tx.Commit(ctx)
	return nil
}

var getNovel = `
SELECT FROM novel WHERE id = $1;
`
