package api_test

import (
	"hlaing-backend-test-1/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserWithEmptyBody(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	CallTestAPI("POST", "/register", nil, &out, headers)
	assert.False(t, out.Success)
	assert.Equal(t, out.Message, "empty request body")
}

func TestRegisterUserWithInvalidParams(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	var params = map[string]string{"name": "thura", "email": "t@random.org"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.False(t, out.Success)
	assert.Contains(t, out.Message, "Field validation for 'Password' failed on the 'required' tag")

	params = map[string]string{"name": "thura", "email": "tre", "password": "letmein"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.False(t, out.Success)
	assert.Contains(t, out.Message, "Field validation for 'Email' failed on the 'email' tag")
}

func TestRegisterUserWithValidParams(t *testing.T) {
	var out = api.DefaultResponse{}
	var headers = map[string]string{}

	var params = map[string]string{"name": "thura", "email": "t@o.com", "password": "letmein"}
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.True(t, out.Success)
	assert.Contains(t, out.Message, "user created successfully")

	// duplicate user should fail
	CallTestAPI("POST", "/register", params, &out, headers)
	assert.False(t, out.Success)
	assert.Contains(t, out.Message, "Error 1062: Duplicate entry")
}
