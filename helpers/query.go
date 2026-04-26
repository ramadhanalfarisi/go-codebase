package helpers

import (
	"database/sql"
)

type QueryHelperInterface interface {
	Insert(sql string, args []any, params ...any) error
	Select(sql string, args []any, params ...any) error
	SelectRow(sql string, args []any, params ...any) error
	Update(sql string, args []any, params ...any) error
	Delete(sql string, args []any) error
}

type QueryHelper struct {
	DB *sql.DB
}

func NewQueryHelper(db *sql.DB) QueryHelperInterface {
	return &QueryHelper{DB: db}
}

func (q *QueryHelper) Insert(sql string, args []any, params ...any) error {
	row := q.DB.QueryRow(sql, args...)
	return row.Scan(params...)
}

func (q *QueryHelper) Select(sql string, args []any, params ...any) error {
	rows, err := q.DB.Query(sql, args...)
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

func (q *QueryHelper) SelectRow(sql string, args []any, params ...any) error {
	row := q.DB.QueryRow(sql, args...)
	return row.Scan(params...)
}

func (q *QueryHelper) Update(sql string, args []any, params ...any) error {
	row := q.DB.QueryRow(sql, args...)
	return row.Scan(params...)
}

func (q *QueryHelper) Delete(sql string, args []any) error {
	_, execErr := q.DB.Exec(sql, args...)
	if execErr != nil {
		Error(execErr)
		return execErr
	}
	return nil
}
