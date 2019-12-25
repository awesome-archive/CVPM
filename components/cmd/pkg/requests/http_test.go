package requests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	c := Client{Endpoint: ""}
	resp := c.Get("http://httpbin.org/get", nil)
	assert.NotEmpty(t, resp.String())
}

func TestPost(t *testing.T) {
	c := Client{Endpoint: ""}
	resp := c.Post("http://httpbin.org/post", map[string]string{"One": "Two"})
	assert.NotEmpty(t, resp.String())
}
