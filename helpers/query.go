package helpers

import "database/sql"

func Insert(db *sql.DB, sql string, args ...any) error {
	tx, err := db.Begin()
	stmt, err := tx.Prepare(sql)
	if err != nil {
		Error(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(args); err != nil {
		Error(err)
		return err
	} else {
		return nil
	}
}

func Select(db *sql.DB, sql string) *sql.Rows {
	rows, err := db.Query(sql, nil)
	if err != nil {
		Error(err)
		return nil
	} else {
		return rows
	}
}
