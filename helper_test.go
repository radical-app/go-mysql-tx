package mysql_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/radical-app/go-mysql-tx"
	"os"
	"testing"
)

func CreateDB(t *testing.T) (db *sql.DB, ctx context.Context) {

	os.Setenv("TEST_DB_HOST","localhost")
	os.Setenv("TEST_DB_NAME","test_db")
	os.Setenv("TEST_DB_PASSWORD","qweuidewfnjewf8")
	os.Setenv("TEST_DB_PORT","3306")
	os.Setenv("TEST_DB_USER","root")

	ctx = context.Background()

	c := mysql.ConfigFromEnvs("TEST")
	db, err := mysql.Open(c, ctx)
	assert.Nil(t, err)
	q := TEST_TABLE_CREATE
	tx, err := mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	_, err = mysql.TxPush(tx, ctx, q)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	return db, ctx
}

func DestroyDB(t *testing.T, db *sql.DB, ctx context.Context) *sql.DB {

	tx, err := mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	_, err = mysql.TxPush(tx, ctx, TEST_TABLE_DROP)
	assert.Nil(t, err)
	assert.Nil(t, tx.Commit())

	return db
}
