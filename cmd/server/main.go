package main

import (
	"github.com/xr1337/gin_bootstrap/gin"
	"github.com/xr1337/gin_bootstrap/mem"
)

func main() {

	quoteService := mem.NewQuoteService()
	quoteController := gin.NewQuoteController(quoteService)
	logincController := gin.NewLoginController(nil)
	server := gin.NewServerDefault(quoteController, logincController)
	server.Run()
}
