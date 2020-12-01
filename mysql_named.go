package mysql

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Named map[string]interface{}

//Query Execution on Transaction
func TxFetchRows(tx *sql.Tx, ctx context.Context, query string, args map[string]interface{}) (rows *sql.Rows, err error) {

	if tx == nil {
		return rows, errors.New("Empty transaction")
	}

	qqm, ordArgs, err  := NamedParameters(query, args)
	if err != nil {
		return rows, err
	}

	rows, err = tx.QueryContext(ctx, qqm, ordArgs...)
	if err != nil {
		return rows, err
	}

	return rows, err
}

//Query Execution on Transaction
func TxPush(tx *sql.Tx, ctx context.Context, query string, args map[string]interface{}) (inc int64, err error) {

	if tx == nil {
		return inc, errors.New("Empty transaction")
	}
	qqm, ordArgs, err  := NamedParameters(query, args)
	if err != nil {
		return inc, err
	}

	res, err := tx.ExecContext(ctx, qqm, ordArgs...)
	if err != nil {
		return inc, err
	}
	inc, err = res.LastInsertId()

	return inc, err
}


// FetchRowsPrepared Is like TxCreate+TxFetchRow Remember to run  Tx.Commit() or Tx.Rollback()
func FetchRows(db *sql.DB, ctx context.Context, query string, args map[string]interface{}) (rows *sql.Rows, tx *sql.Tx, err error) {

	tx, err = TxCreate(db, ctx)
	if err!=nil {
		return rows, tx, err
	}

	rows, err = TxFetchRows(tx, ctx, query, args)
	if err != nil {
		return rows, tx , err
	}

	return rows, tx, err
}

// PushPrepared Is like  TxCreate+TxPush+tx.Commit  in once
func Push(db *sql.DB, ctx context.Context, query string, args map[string]interface{}) (inc int64, err error) {

	tx, err := TxCreate(db, ctx)
	if err!=nil {
		return inc, err
	}
	inc, err = TxPush(tx, ctx, query, args)
	if err != nil {
		return inc, err
	}

	err = tx.Commit()

	return inc, err
}
