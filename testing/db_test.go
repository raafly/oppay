package testing

import (
	"testing"

	"github.com/raafly/ewallet/config"
	"github.com/raafly/ewallet/db"

	"github.com/stretchr/testify/assert"
)

var cfg = config.NewConfig()

func TestDB(t *testing.T) {
	db := db.NewPostgres(cfg)
	assert.Nil(t, db)
}