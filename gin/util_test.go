package gin_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toJSONReader(q interface{}) io.Reader {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(q)
	return b
}

func toObject(t *testing.T, w *httptest.ResponseRecorder, obj interface{}) {
	t.Helper()
	data := w.Body.String()
	err := json.Unmarshal([]byte(data), &obj)
	assert.NoError(t, err)
}

func assertJSONHeaderResponse(t *testing.T, w *httptest.ResponseRecorder) {
	t.Helper()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

func performRequest(t *testing.T, r http.Handler, method, path string, reader io.Reader, cookie string) *httptest.ResponseRecorder {
	t.Helper()
	req, err := http.NewRequest(method, path, reader)
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
