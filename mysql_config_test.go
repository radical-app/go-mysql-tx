package toolbox_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/radical-app/sql-fast-toolbox"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestConfig_GetConnection(t *testing.T) {

	prefix := strings.ToUpper("TEST_"+t.Name()+"_")
	err := os.Setenv(prefix + "DB_USER", "USER")
	assert.Nil(t, err)
	err = os.Setenv(prefix + "DB_PASSWORD", "PASSWORD")
	assert.Nil(t, err)
	err = os.Setenv(prefix + "DB_NAME", "NAME")
	assert.Nil(t, err)
	err = os.Setenv(prefix + "DB_HOST", "HOST")
	assert.Nil(t, err)
	err = os.Setenv(prefix + "DB_PORT", "3306")
	assert.Nil(t, err)

	c := toolbox.ConfigFromEnvs(prefix)

	assert.Equal(t, "USER:PASSWORD@tcp(HOST:3306)/NAME?parseTime=true&multiStatements=true&loc=UTC&charset=utf8", c.GetConnection())
}
