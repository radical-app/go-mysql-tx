package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Open(cnf ConfigOpenable, ctx context.Context) (*sql.DB, error) {

	db, err := sql.Open("mysql", cnf.GetConnection())
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}

func SetMaxLifetimeMins(db *sql.DB, minutesMaxLifeConn time.Duration) {
	db.SetConnMaxLifetime(minutesMaxLifeConn*time.Minute)
}

//Create a Transaction
func TxCreate(db *sql.DB, ctx context.Context) (*sql.Tx , error) {
	return db.BeginTx(ctx, nil)
}

//Query Execution on Transaction
func TxFetchRows(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {

	if tx == nil {
		return rows, errors.New("Empty transaction")
	}

	rows, err = tx.QueryContext(ctx, query, args...)
	if err != nil {
		return rows, err
	}

	return rows, err
}

//Query Execution on Transaction
func TxPush(tx *sql.Tx, ctx context.Context, query string, args ...interface{}) (inc int64, err error) {

	if tx == nil {
		return inc, errors.New("Empty transaction")
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return inc, err
	}
	inc, err = res.LastInsertId()

	return inc, err
}

// IsErrorRollback should replace the err!=nil check on Error it Rollback the tx
func IsErrorRollback(err error, tx *sql.Tx) bool {
	if err != nil {
		_ = tx.Rollback()
		return true
	}

	return false
}

// FetchRows Is like TxCreate+TxFetchRows+tx.Commit in once
func FetchRows(db *sql.DB, ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {

	tx, err := TxCreate(db, ctx)
	if err!=nil {
		return rows, err
	}
	rows, err = TxFetchRows(tx, ctx, query, args...)
	fmt.Print(err)
	if err != nil {
		return rows, err
	}

	err = tx.Commit()

	return rows, err
}

// Push Is like  TxCreate+TxPush+tx.Commit  in once
func Push(db *sql.DB, ctx context.Context, query string, args ...interface{}) (inc int64, err error) {

	tx, err := TxCreate(db, ctx)
	if err!=nil {
		return inc, err
	}
	inc, err = TxPush(tx, ctx, query, args...)
	if err != nil {
		return inc, err
	}

	err = tx.Commit()

	return inc, err
}
