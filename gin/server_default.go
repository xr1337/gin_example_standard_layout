package gin

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	app "github.com/xr1337/gin_bootstrap"
)

// Ensure ServerDefault implements app.Server
var _ app.Server = &ServerDefault{}

type ServerDefault struct {
	quoteController *QuoteController
	loginController *LoginController
	router          *gin.Engine
}

func NewServerDefault(qc *QuoteController, lc *LoginController) *ServerDefault {
	server := &ServerDefault{quoteController: qc, loginController: lc}
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.GET("/list", server.quoteController.List)
	r.GET("/quote/:id", server.quoteController.Get)
	r.POST("/quote", server.quoteController.Add)

	r.POST("/user/login", server.loginController.Login)
	private := r.Group("/")
	private.Use(server.loginController.AuthRequired)
	{
		private.GET("/user/me", server.loginController.Me)
	}

	server.router = r
	return server
}

func (s *ServerDefault) Router() http.Handler {
	return s.router
}

func (s *ServerDefault) QuoteController() app.QuoteController {
	return s.quoteController
}

func (s *ServerDefault) LoginController() app.LoginController {
	return s.loginController
}

func (s *ServerDefault) Run() {
	s.router.Run()
}
