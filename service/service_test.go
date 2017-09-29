package service

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	assert := assert.New(t)

	helloHandler := &HelloHandler{}
	server := httptest.NewServer(helloHandler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	assert.NoError(err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)

	assert.Equal(200, resp.StatusCode)
	assert.Equal("text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal("Hello.", string(body))
}

func TestHiHandler(t *testing.T) {
	assert := assert.New(t)

	hiHandler := &HiHandler{}
	server := httptest.NewServer(hiHandler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	assert.NoError(err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)

	assert.Equal(200, resp.StatusCode)
	assert.Equal("text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal("Hi.", string(body))
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	handlers := Handlers()

	ts := httptest.NewServer(handlers)
	defer ts.Close()

	f := func(path, expected string) {
		res, err := http.Get(ts.URL + path)
		assert.NoError(err)
		byts, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		res.Body.Close()
		assert.Equal(expected, string(byts))
	}
	f("/hello", "Hello.")
	f("/hi", "Hi.")
}
