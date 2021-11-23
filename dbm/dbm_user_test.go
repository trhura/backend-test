package dbm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserManager(t *testing.T) {
	// Check validations
	err := dbManager.AddNewUser("", "thura@random.org", "letmein")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name must not be empty")

	err = dbManager.AddNewUser("thura", "", "letmein")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email must not be empty")

	err = dbManager.AddNewUser("thura", "thura", "letmein")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mail: missing '@'")

	err = dbManager.AddNewUser("thura", "thura@random.org", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "password must not be empty")

	// it is successful all parameters are good
	err = dbManager.AddNewUser("thura", "thura@random.org", "letmein")
	assert.NoError(t, err)

	// it must fail when there is an existing email address in the system
	err = dbManager.AddNewUser("thura", "thura@random.org", "letmein")
	assert.Error(t, err)

	// Check get user with invalid user email
	user, err := dbManager.GetUser("asdf@random.org")
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sql: no rows in result set")

	// Check get user with valid user email
	user, err = dbManager.GetUser("thura@random.org")
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.EqualValues(t, user.Email, "thura@random.org")
	assert.EqualValues(t, user.Name, "thura")
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)

	// Authenticate with invalid password
	user, err = dbManager.AuthenticateUser("thura@random.org", "asdf")
	assert.NotEmpty(t, user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not the hash of the given password")

	// Authenticate with valid password
	user, err = dbManager.AuthenticateUser("thura@random.org", "letmein")
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.EqualValues(t, user.Email, "thura@random.org")
	assert.EqualValues(t, user.Name, "thura")
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)

}
