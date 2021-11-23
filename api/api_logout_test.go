package api_test

import (
	"hlaing-backend-test-1/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogoutUserSuccessful(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	var params = map[string]string{"name": "usere", "email": "e@m.com", "password": "pass"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.True(t, out.Success)

	headers = map[string]string{
		"Authorization": "Basic " + basicAuth("e@m.com", "pass"),
	}
	resp := CallTestAPI("POST", "/login", nil, &out, headers)
	assert.True(t, out.Success)
	assert.Contains(t, out.Message, "user login successful")

	cookies := resp.Cookies()
	assert.Len(t, cookies, 1)
	headers = map[string]string{"Cookie": cookies[0].String()}

	resp = CallTestAPI("POST", "/logout", nil, &out, headers)
	assert.True(t, out.Success)
	assert.Contains(t, out.Message, "user logout successful")

	cookies = resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "ssn")
	assert.Empty(t, cookies[0].Value)
	assert.NotEmpty(t, cookies[0].Path, "/")
	assert.EqualValues(t, cookies[0].MaxAge, -1)
	assert.LessOrEqual(t, cookies[0].Expires.Unix(), time.Now().Unix())
}
