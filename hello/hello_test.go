package hello

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handleHello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(200, resp.StatusCode)
	assert.Equal("text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal("Hello.", string(body))
}

func TestHiHandler(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handleHi(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(200, resp.StatusCode)
	assert.Equal("text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal("Hi.", string(body))
}
