package db

import (
	"testing"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/stretchr/testify/assert"
)

func TestStatements(t *testing.T) {
	assert.Nil(t, InitDB())
	_, ok := AuthentificateUser("", "")
	assert.False(t, ok)
	assert.Nil(t, InsertGame(&common.Game{}))
	us, err := QueryUser("search")
	assert.Nil(t, err)
	assert.Zero(t, len(us))
}
