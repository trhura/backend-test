package api_test

import (
	"encoding/base64"
	"hlaing-backend-test-1/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

// copied from golang http package
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func TestLoginuserWithoutHeader(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	CallTestAPI("POST", "/login", nil, &out, headers)
	assert.False(t, out.Success)
	assert.Equal(t, out.Message, "unable to parse basic auth")
}

func TestLoginuserWithInvalidHeader(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{"Authorization": "Basic " + basicAuth("a@m.com", "x")}

	// no such user
	CallTestAPI("POST", "/login", nil, &out, headers)
	assert.False(t, out.Success)
	assert.Contains(t, out.Message, "sql: no rows in result set")

	var params = map[string]string{"name": "usera", "email": "a@m.com", "password": "pass"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.True(t, out.Success)

	// wrong password
	CallTestAPI("POST", "/login", nil, &out, headers)
	assert.False(t, out.Success)
	assert.Contains(t, out.Message, "not the hash of the given password")
}

func TestLoginUserSuccessful(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	var params = map[string]string{"name": "userb", "email": "b@m.com", "password": "pass"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.True(t, out.Success)

	headers = map[string]string{
		"Authorization": "Basic " + basicAuth("b@m.com", "pass"),
	}
	resp := CallTestAPI("POST", "/login", nil, &out, headers)
	assert.True(t, out.Success)
	assert.Contains(t, out.Message, "user login successful")

	cookies := resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "ssn")
	assert.NotEmpty(t, cookies[0].Value)
	assert.NotEmpty(t, cookies[0].Path, "/")
	assert.EqualValues(t, cookies[0].MaxAge, int(dbManager.Envars.SessionLifetime.Seconds()))
}
