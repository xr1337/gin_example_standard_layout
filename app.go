package app

import (
	"net/http"
	"time"
)

type Quote struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Email string
type User struct {
	Email     Email     `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type QuoteService interface {
	CreateOrUpdate(quote Quote) (int, error)
	GetQuote(id int) (Quote, error)
	ListQuotes() []Quote
}

type QuoteController interface {
	QuoteService() QuoteService
}

type LoginService interface {
	UserExist(email Email) bool
}

type LoginController interface {
	LoginService() LoginService
}

type Server interface {
	QuoteController() QuoteController
	LoginController() LoginController
	Router() http.Handler
	Run()
}
