package mysql_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/radical-app/go-mysql-tx"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestOpenCommitRollbackSingle(t *testing.T) {
	db, ctx := CreateDB(t)
	//-----

	tx, err := mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	firstInc, err := mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "inserting data multiple commit")
	assert.Nil(t, err)
	assert.True(t, firstInc > 0)
	assert.Nil(t, tx.Commit())
	//-----

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	inc, err := mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "inserting second data on same transaction and commit")
	assert.Nil(t, err)
	assert.True(t, inc > 0)
	assert.Nil(t, tx.Commit())

	//-----

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	rows, err := mysql.TxFetchRowsPrepared(tx, ctx, TEST_TABLE_SELECT, firstInc)
	assert.Nil(t, err)
	assert.Nil(t, rows.Err())

	count := 0
	for rows.Next() {
		var id int64= 0
		name := ""
		count++

		err = rows.Scan(&id, &name)
		assert.Nil(t, err)
		assert.Equal(t, firstInc, id)
		assert.True(t, name != "")
	}
	assert.True(t, count > 0)

	assert.Nil(t, rows.Close())
	assert.Nil(t, tx.Commit())

	DestroyDB(t, db, ctx)
}



func TestOpenCommitRollback(t *testing.T) {
	db, ctx := CreateDB(t)
	//-----

	tx, err := mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	firstInc, err := mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "inserting data multiple commit")
	assert.Nil(t, err)
	assert.True(t, firstInc > 0)
	assert.Nil(t, tx.Commit())
	//-----

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	inc, err := mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "inserting second data on same transaction and commit")
	assert.Nil(t, err)
	assert.True(t, inc > 0)
	assert.Nil(t, tx.Commit())

	//-----

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	inc, err = mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "Third insert on committed tx")
	assert.Nil(t, err)
	assert.True(t, inc > 0)
	assert.Nil(t, tx.Commit())

	//-----

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	inc, err = mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "Fourth insert on committed tx then rollback and try to insert again")
	assert.Nil(t, err)
	assert.True(t, inc > 0)
	assert.Nil(t, tx.Rollback())

	inc, err = mysql.TxPushPrepared(tx, ctx, TEST_TABLE_INSERT, "Should fail has been already roll-backed")
	assert.NotNil(t, tx.Commit())

	//----- Select

	tx, err = mysql.TxCreate(db, ctx)
	assert.Nil(t, err)
	rows, err := mysql.TxFetchRowsPrepared(tx, ctx, TEST_TABLE_SELECT, firstInc)
	assert.Nil(t, err)
	assert.Nil(t, rows.Err())

	count := 0
	for rows.Next() {
		var id int64= 0
		name := ""
		count++

		err = rows.Scan(&id, &name)
		assert.Nil(t, err)
		assert.Equal(t, firstInc, id)
		assert.True(t, name != "")
	}
	assert.True(t, count > 0)

	assert.Nil(t, rows.Close())
	assert.Nil(t, tx.Commit())

	DestroyDB(t, db, ctx)
}
