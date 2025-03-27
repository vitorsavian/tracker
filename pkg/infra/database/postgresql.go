package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/tracker/pkg/adapter"
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

func (p *PSQL) CreateNovel(novel *adapter.CreateNovelDatabaseRequestAdapter) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin create novel transction: %v\n", err)
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

func (p *PSQL) UpdateNovel(novel *adapter.UpdateNovelDatabaseRequestAdapter) error {
	ctx := context.Background()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("Unable to begin for update novel transction: %v\n", err)
		return err
	}

	if _, err = tx.Exec(ctx, updateNovel,
		novel.Id, novel.Name, novel.Page, novel.Finished, utils.UTCTime()); err != nil {
		logrus.Errorf("Unable to exec query to update novel: %v\n", err)

		if err = tx.Rollback(ctx); err != nil {
			logrus.Errorf("Unable to rollback update novel transaction: %v\n", err)
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
		logrus.Errorf("Unable to begin for delete novel transction: %v\n", err)
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
  SELECT novel.id::text, novel.name, novel.page, novel.finished, novel.created_at, novel.updated_at FROM novel WHERE id = $1;
`

func (p *PSQL) GetNovel(id string) (*adapter.GetNovelDatabaseResponseAdapter, error) {
	ctx := context.Background()

	novel := adapter.GetNovelDatabaseResponseAdapter{}

	err := p.conn.QueryRow(ctx, getNovel, id).Scan(&novel.Id, &novel.Name,
		&novel.Page, &novel.Finished,
		&novel.CreatedAt, &novel.UpdatedAt)
	if err == pgx.ErrNoRows {
		logrus.Errorf("Unable to get any rows from the get query: %v\n", err)
		return nil, err
	} else if err != nil {
		logrus.Errorf("Unable to begin for delete novel transction: %v\n", err)
		return nil, err
	}

	return &novel, nil
}

var getAllNovels = `
SELECT novel.id, novel.name, novel.page, novel.finished, novel.created_at, novel_updated_at FROM novel
`

// func (p *PSQL) GetNovel(id string) error {
// 	ctx := context.Background()
// 	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		logrus.Errorf("Unable to begin for delete novel transction: %v\n", err)
// 		return err
// 	}
//
// 	if _, err = tx.Exec(ctx, deleteNovel, id); err != nil {
// 		logrus.Errorf("Unable to exec query to delete novel: %v\n", err)
//
// 		if err = tx.Rollback(ctx); err != nil {
// 			logrus.Errorf("Unable to rollback delete novel transaction: %v\n", err)
// 			return err
// 		}
//
// 		return err
// 	}
//
// 	tx.Commit(ctx)
// 	return nil
// }
