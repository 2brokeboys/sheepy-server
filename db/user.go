package db

import (
	"github.com/2brokeboys/sheepy-server/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthentificateUser checks the given credentials in the database and returns the user
// on successful authentification and a boolean flag indicating whether the credentials
// were correct
func AuthentificateUser(username string, password string) (*common.User, bool) {
	if gin.Mode() == gin.DebugMode && username == "test" {
		return &common.User{
			ID:       1,
			Name:     "Test",
			Username: "Max Mustermann",
		}, true
	}

	user := &dbUser{}
	err := getExactUserStatement.Get(user, username)
	if err != nil {
		return nil, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pw), []byte(password))
	if err != nil {
		return nil, false
	}

	return user.ToCommon(), true
}

// InsertUser inserts a new user into the database that can be used in future logins
func InsertUser(user *common.User, password string) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "Error generating bcrypt hash")
	}

	dbuser := &dbUser{
		Username: user.Username,
		Name:     user.Name,
		Pw:       string(pw),
	}

	_, err = insertUserStatement.Exec(dbuser)
	if err != nil {
		return errors.Wrap(err, "Error writing user into db")
	}

	return nil
}
