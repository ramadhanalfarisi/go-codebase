package helpers

import "database/sql"

func Insert(db *sql.DB, sql string, args []any) error {
	stmt, err := db.Prepare(sql)
	if err != nil {
		Error(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(args...); err != nil {
		Error(err)
		return err
	} else {
		return nil
	}
}

func Select(db *sql.DB, sql string, args []any) *sql.Rows {
	rows, err := db.Query(sql, args...)
	if err != nil {
		Error(err)
		return nil
	} else {
		return rows
	}
}

func SelectRow(db *sql.DB, sql string, args []any) *sql.Row {
	row := db.QueryRow(sql, args...)
	return row
}
