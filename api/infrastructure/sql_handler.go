package infrastructure

import (
	"api/interface/repository"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func (handler SqlHandler) Execute(statement string, args ...interface{}) (repository.Result, error) {
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return nil, err
	}
	res := SqlResult{Result: result}
	return res, nil
}

func (handler SqlHandler) Query(statement string, args ...interface{}) (repository.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	row := SqlRow{Rows: rows}
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsEffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

func NewSqlHandler() repository.SqlHandler {
	conn, err := sql.Open("mysql", "recipe:password@tcp(db:3306)/recipe_db")
	if err != nil {
		log.Fatal(err.Error)
	}
	sqlHandler := SqlHandler{Conn: conn}
	return sqlHandler
}
