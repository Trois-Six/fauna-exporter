package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_Index(t *testing.T) {
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	New("/metrics").Index(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)

	body, err := ioutil.ReadAll(rw.Body)
	require.NoError(t, err)
	assert.Contains(t, string(body), "<p><a href='/metrics'>Metrics</a></p>")
}

func TestHandlers_OK(t *testing.T) {
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	New("/metrics").OK(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)

	body, err := ioutil.ReadAll(rw.Body)
	require.NoError(t, err)
	assert.Equal(t, string(body), "OK")
}
