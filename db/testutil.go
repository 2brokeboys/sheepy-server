package db

import (
	"testing"

	"github.com/2brokeboys/sheepy-server/common"
	"github.com/stretchr/testify/assert"
)

// AddTestUsers creates some users for unit testing
func AddTestUsers(t *testing.T) {
	// Add some test users
	assert.Nil(t, InsertUser(&common.User{
		Username: "foo",
	}, "123456"))
	assert.Nil(t, InsertUser(&common.User{
		Username: "bar",
	}, "123456"))
	assert.Nil(t, InsertUser(&common.User{
		Username: "moo",
	}, "123456"))
}
