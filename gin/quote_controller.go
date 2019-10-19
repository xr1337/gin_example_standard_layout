package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/xr1337/gin_bootstrap"
)

var _ app.QuoteController = &QuoteController{}

type QuoteAdd struct {
	Text string `json:"text" binding:"required"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

type QuoteController struct {
	quoteService app.QuoteService
}

func NewQuoteController(service app.QuoteService) *QuoteController {
	return &QuoteController{quoteService: service}
}

func (p *QuoteController) QuoteService() app.QuoteService {
	return p.quoteService
}

func (p *QuoteController) Get(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{app.QuoteInvalidIdErr.Error()})
		return
	}
	quote, err := p.quoteService.GetQuote(int_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{app.QuoteInvalidIdErr.Error()})
		return
	}
	c.JSON(200, quote)
}

func (p *QuoteController) Add(c *gin.Context) {
	var quote QuoteAdd
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{app.QuoteInvalidFormatErr.Error()})
		return
	}
	saveQuoteObj := app.Quote{Text: quote.Text}
	id, err := p.quoteService.CreateOrUpdate(saveQuoteObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{app.QuoteInvalidFormatErr.Error()})
		return
	}
	saveQuoteObj.Id = id
	c.JSON(200, saveQuoteObj)
}

func (p *QuoteController) List(c *gin.Context) {
	quotes := p.quoteService.ListQuotes()
	c.JSON(200, quotes)
}
