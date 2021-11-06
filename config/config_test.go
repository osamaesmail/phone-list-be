package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert.NotEmpty(t, Cfg().AppMode, "APP_MODE")
	assert.NotZero(t, Cfg().AppPort, "APP_PORT")
	assert.NotZero(t, Cfg().AppDefaultPageSize, "APP_DEFAULT_PAGE_SIZE")
	assert.NotZero(t, Cfg().AppMaxPageSize, "APP_MAX_PAGE_SIZE")
	assert.NotEmpty(t, Cfg().DBDriver, "DB_DRIVER")
	assert.NotEmpty(t, Cfg().Sqlite3Path, "SQLITE3_PATH")
}
