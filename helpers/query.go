package helpers

import (
	"context"
	"database/sql"
	"strings"
)

type QueryHelperInterface interface {
	Insert(ctx context.Context, sql string, args []any, params ...any) error
	Select(ctx context.Context, sql string, args []any, params ...any) error
	SelectRow(ctx context.Context, sql string, args []any, params ...any) error
	Update(ctx context.Context, sql string, args []any, params ...any) error
	Delete(ctx context.Context, sql string, args []any) error
}

type QueryHelper struct {
	DB *sql.DB
}

func NewQueryHelper(db *sql.DB) QueryHelperInterface {
	return &QueryHelper{DB: db}
}

func (q *QueryHelper) Insert(ctx context.Context, sql string, args []any, params ...any) error {
	err := q.updateInsert(ctx, sql, args, params...)
	return err
}

func (q *QueryHelper) Select(ctx context.Context, sql string, args []any, params ...any) error {
	rows, err := q.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		Error(err)
		return err
	} else {
		for rows.Next() {
			if err := rows.Scan(params...); err != nil {
				Error(err)
				return err
			}
		}
		return nil
	}
}

func (q *QueryHelper) SelectRow(ctx context.Context, sql string, args []any, params ...any) error {
	row := q.DB.QueryRowContext(ctx, sql, args...)
	return row.Scan(params...)
}

func (q *QueryHelper) Update(ctx context.Context, sql string, args []any, params ...any) error {
	err := q.updateInsert(ctx, sql, args, params...)
	return err
}

func (q *QueryHelper) Delete(ctx context.Context, sql string, args []any) error {
	_, execErr := q.DB.ExecContext(ctx, sql, args...)
	if execErr != nil {
		Error(execErr)
		return execErr
	}
	return nil
}

func (q *QueryHelper) updateInsert(ctx context.Context, sql string, args []any, params ...any) error {
	if strings.Contains(sql, "RETURNING") {
		row := q.DB.QueryRowContext(ctx, sql, args...)
		return row.Scan(params...)
	}
	_, execErr := q.DB.ExecContext(ctx, sql, args...)
	if execErr != nil {
		Error(execErr)
		return execErr
	}
	return nil
}
