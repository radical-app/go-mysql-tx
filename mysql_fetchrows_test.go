package mysql_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/radical-app/go-mysql-tx"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestPushAndFetchRows(t *testing.T) {

	db, ctx := CreateDB(t)
	//-----
	i, err := mysql.PushPrepared(db, ctx, TEST_TABLE_INSERT, "new insert")
	assert.Nil(t, err)
	assert.True(t, i>0)


	//-----
	rows, tx, err := mysql.FetchRowsPrepared(db, ctx, TEST_TABLE_SELECT, i)
	assert.Nil(t, err)
	count := 0
	for rows.Next() {
		var id int64= 0
		name := ""
		count++

		err = rows.Scan(&id, &name)
		assert.Nil(t, err)
		assert.Equal(t, i, id)
		assert.True(t, name != "")
	}
	assert.True(t, count > 0)
	assert.Nil(t,tx.Commit())
	// ---
	DestroyDB(t, db, ctx)
}
