package healthcheck

import (
	"io/ioutil"
	"testing"

	"github.com/ampproject/amppackager/packager/mux"
	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckOK(t *testing.T) {
	handler, err := New()
	require.NoError(t, err)

	resp := pkgt.Get(t, mux.New(nil, nil, nil, handler), "/healthcheck")
	defer resp.Body.Close()
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "no-store", resp.Header.Get("Cache-Control"))
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	expected := string(`{
	"cert": "OK",
	"http": "OK",
	"rtv": "OK"
}`)
	assert.Equal(t, expected, string(body))
}

func TestHealthCheckRtvError(t *testing.T) {
	handler, err := New()
	require.NoError(t, err)

	handler.rtv = "RTV ERROR: Fetch Failed"

	resp := pkgt.Get(t, mux.New(nil, nil, nil, handler), "/healthcheck")
	defer resp.Body.Close()
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "no-store", resp.Header.Get("Cache-Control"))
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, 500, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	expected := string(`{
	"cert": "OK",
	"http": "OK",
	"rtv": "RTV ERROR: Fetch Failed"
}`)
	assert.Equal(t, expected, string(body))
}
