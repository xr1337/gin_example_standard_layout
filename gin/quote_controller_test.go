package gin_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	app "github.com/xr1337/gin_bootstrap"
	"github.com/xr1337/gin_bootstrap/gin"
	"github.com/xr1337/gin_bootstrap/mem"
)

func TestListQuotes(t *testing.T) {
	t.Run("Test with 2 items", func(t *testing.T) {
		server := emptyServer()
		want := map[string]int{
			"pink floyd": 1,
			"john mason": 2,
		}
		for k := range want {
			createQuote(t, server.Router(), app.Quote{Text: k})
		}

		w := performRequest(t, server.Router(), "GET", "/list", nil, "")
		assertJSONHeaderResponse(t, w)

		var got []app.Quote
		toObject(t, w, &got)
		for _, item := range got {
			assert.NotNil(t, want[item.Text])
		}
	})
}

func TestGetOneQuote(t *testing.T) {
	server := emptyServer()
	w := createQuote(t, server.Router(), app.Quote{Text: "power"})
	var ot app.Quote
	toObject(t, w, &ot)
	quoteID := ot.Id

	t.Run("test getting one item", func(t *testing.T) {
		w := performRequest(t, server.Router(), "GET", fmt.Sprintf("/quote/%d", quoteID), nil, "")
		assertJSONHeaderResponse(t, w)

		var got app.Quote
		toObject(t, w, &got)
		assert.Equal(t, "power", got.Text)
		assert.Equal(t, quoteID, got.Id)
	})
	t.Run("test bad input", func(t *testing.T) {
		badTest := []string{
			"/quote/somevalue",
			"/quote/123",
		}
		for _, val := range badTest {
			w := performRequest(t, server.Router(), "GET", val, nil, "")
			assert.Equal(t, w.Code, http.StatusBadRequest)

			var got map[string]string
			toObject(t, w, &got)
			assert.Equal(t, got["error"], app.QuoteInvalidIdErr.Error())
		}
	})
}

func TestSave(t *testing.T) {

	t.Run("create new quote", func(t *testing.T) {
		server := emptyServer()
		want := app.Quote{Text: "gool"}
		w := createQuote(t, server.Router(), want)
		var got app.Quote
		toObject(t, w, &got)
		assert.Equal(t, got.Text, want.Text)
		assert.Equal(t, 1, got.Id)
	})

	t.Run("create new quote ignore id spoof", func(t *testing.T) {
		server := emptyServer()
		want := app.Quote{Id: 3, Text: "gool"}
		w := createQuote(t, server.Router(), want)
		var got app.Quote
		toObject(t, w, &got)
		assert.Equal(t, got.Text, want.Text)
		assert.Equal(t, 1, got.Id)
	})
	t.Run("create empty quote should fail", func(t *testing.T) {
		server := emptyServer()
		want := app.Quote{Text: ""}
		b := toJSONReader(want)
		w := performRequest(t, server.Router(), "POST", "/quote", b, "")

		var got map[string]string
		toObject(t, w, &got)
		assert.Equal(t, got["error"], app.QuoteInvalidFormatErr.Error())
	})
}

func createQuote(t *testing.T, handler http.Handler, q app.Quote) *httptest.ResponseRecorder {
	t.Helper()
	b := toJSONReader(q)
	w := performRequest(t, handler, "POST", "/quote", b, "")
	assertJSONHeaderResponse(t, w)
	return w
}

func emptyServer() (server *gin.ServerDefault) {
	store := mem.NewQuoteService()
	sut := gin.NewQuoteController(store)
	server = gin.NewServerDefault(sut, nil)
	return
}
