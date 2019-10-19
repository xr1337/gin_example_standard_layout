package mem

import (
	"time"

	app "github.com/xr1337/gin_bootstrap"
)

// ensure that this conforms to interface
var _ app.QuoteService = &QuoteService{}

type QuoteService struct {
	items map[int]app.Quote
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		items: make(map[int]app.Quote),
	}
}

func isNewQuote(q app.Quote) bool {
	return q.Id <= 0
}

func (q *QuoteService) CreateOrUpdate(quote app.Quote) (int, error) {
	if quote.Text == "" {
		return 0, app.QuoteMissingErr
	}
	if isNewQuote(quote) {
		quote.CreatedAt = time.Now()
		quote.UpdatedAt = quote.CreatedAt
		quote.Id = q.maxId() + 1

		q.items[quote.Id] = quote
	} else {
		q.updateQuote(quote)
	}
	return quote.Id, nil
}

func (i *QuoteService) GetQuote(id int) (app.Quote, error) {
	var val app.Quote
	if val, ok := i.items[id]; ok {
		return val, nil
	}
	return val, app.QuoteInvalidIdErr
}

func (i *QuoteService) ListQuotes() []app.Quote {
	list := make([]app.Quote, 0, len(i.items))
	for _, val := range i.items {
		list = append(list, val)
	}
	return list
}

func (i *QuoteService) updateQuote(q app.Quote) {
	q.UpdatedAt = time.Now()
	i.items[q.Id] = q
}

func (i *QuoteService) maxId() int {
	max := 0
	for key, _ := range i.items {
		if max < key {
			max = key
		}
	}
	return max
}
