package mem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	app "github.com/xr1337/gin_bootstrap"
	"github.com/xr1337/gin_bootstrap/mem"
)

const ConstQuoteText = "sometext"

func TestSaveQuote(t *testing.T) {
	sut := mem.NewQuoteService()
	q, _ := makeFromText(sut, ConstQuoteText)

	t.Run("test saving one item", func(t *testing.T) {
		assertQuoteHasText(t, sut, q.Id, ConstQuoteText)
	})
	t.Run("test updating item", func(t *testing.T) {
		want := "another one"
		q.Text = want
		sut.CreateOrUpdate(q)
		assertQuoteHasText(t, sut, q.Id, want)
	})
	t.Run("test creating new items", func(t *testing.T) {
		n, err := makeFromText(sut, ConstQuoteText)
		assert.NoError(t, err)
		assert.NotEqual(t, n.Id, q.Id)
	})
	t.Run("save empty quote not allowed", func(t *testing.T) {
		_, err := makeFromText(sut, "")
		assert.EqualError(t, err, app.QuoteMissingErr.Error())
	})
}

func TestGetAllQuote(t *testing.T) {
	sut := mem.NewQuoteService()
	makeFromText(sut, ConstQuoteText)
	makeFromText(sut, "another one")

	got := sut.ListQuotes()
	assert.Equal(t, len(got), 2)
}

func makeFromText(service app.QuoteService, text string) (app.Quote, error) {
	quote := app.Quote{Text: text}
	id, err := service.CreateOrUpdate(quote)
	quote.Id = id
	return quote, err
}

func assertQuoteHasText(t *testing.T, service app.QuoteService, id int, want string) {
	t.Helper()
	got, err := service.GetQuote(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, got.Id)
	assert.Equal(t, want, got.Text)
	assert.NotNil(t, got.CreatedAt)
	assert.NotNil(t, got.UpdatedAt)
}
