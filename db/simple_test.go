package db

import (
	"testing"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/stretchr/testify/assert"
)

func TestStatements(t *testing.T) {
	// Init DB
	assert.Nil(t, InitDB())

	// Empty test db should not contain any users
	_, ok := AuthentificateUser("", "")
	assert.False(t, ok)
	assert.Nil(t, InsertGame(&common.Game{}))
	us, err := QueryUser("search")
	assert.Nil(t, err)
	assert.Zero(t, len(us))

	AddTestUsers(t)

	// Check test user login
	_, ok = AuthentificateUser("foo", "123456")
	assert.True(t, ok)

	// Query "oo" -> should return "foo" and "moo"
	us, err = QueryUser("oo")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(us))
}
