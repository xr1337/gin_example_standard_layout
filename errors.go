package app

const (
	QuoteMissingErr       = Error("Quote is missing")
	QuoteInvalidIdErr     = Error("Quote id is invalid")
	QuoteInvalidFormatErr = Error("Quote is invalid")
)

type Error string

func (q Error) Error() string {
	return string(q)
}
