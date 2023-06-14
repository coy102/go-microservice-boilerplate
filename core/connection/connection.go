package connection

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/v2/util/log"
)

// CrudResult ...
type CrudResult struct {
	InsertedID   int64
	RowsAffected int64
}

type connection struct {
	db *sqlx.DB
}

// Connection ...
type Connection interface {
	ExecuteDBQuery(qapi string, args ...interface{}) ([][]string, error)
	ExecuteDBInsertReturnID(qapi string, args ...interface{}) (CrudResult, error)
	ExecuteDB(qapi string, args ...interface{}) (CrudResult, error)
	DB() *sqlx.DB
	BeginTx() (*sql.Tx, error)
	EndTx(tx *sql.Tx, err error) error
}

func (connection *connection) ExecuteDBQuery(qapi string, args ...interface{}) ([][]string, error) {
	rows, errConnection := connection.db.Query(qapi, args...)

	var records [][]string

	if errConnection == nil {
		records = convertToCsv(rows)
	}

	if rows != nil {
		defer rows.Close()
	}

	return records, errConnection
}

func (connection *connection) ExecuteDB(qapi string, args ...interface{}) (CrudResult, error) {
	result, errConnection := connection.db.Exec(qapi, args...)
	crud := CrudResult{}

	if errConnection != nil {
		log.Info("Failed execute Query", errConnection)
		return crud, errConnection
	}

	crud.RowsAffected, _ = result.RowsAffected()
	return crud, nil
}

func (connection *connection) DB() *sqlx.DB {
	return connection.db
}

func (connection *connection) BeginTx() (*sql.Tx, error) {
	return connection.db.Begin()
}

func (connection *connection) EndTx(tx *sql.Tx, err error) error {
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func convertToCsv(rows *sql.Rows) (results [][]string) {
	cols, err := rows.Columns()
	if err != nil {
		log.Info("Failed to get columns", err)
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	results = append(results, []string{"HEADER"}) //static values on the first index

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		var result []string
		err = rows.Scan(dest...)
		if err != nil {
			log.Info("Failed to scan row", err)
			return
		}

		for _, raw := range rawResult {
			if raw == nil {
				result = append(result, "")
			} else {
				result = append(result, string(raw))
			}
		}

		results = append(results, result)
	}

	return results
}
