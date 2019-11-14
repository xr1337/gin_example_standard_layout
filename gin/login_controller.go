package gin

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	app "github.com/xr1337/gin_bootstrap"
)

type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

const SessionUser = "user"

var _ app.LoginController = &LoginController{}

type LoginController struct {
	loginService app.LoginService
}

func NewLoginController(ls app.LoginService) *LoginController {
	loginController := &LoginController{loginService: ls}
	return loginController
}

func (l *LoginController) LoginService() app.LoginService {
	return l.loginService
}

func (l *LoginController) Me(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}

func (l *LoginController) Login(c *gin.Context) {
	var login Login

	if err := c.ShouldBind(&login); err != nil {
		return
	}

	if login.Email == "john@gmail.com" && login.Password == "good" {
		session := sessions.Default(c)
		session.Set(SessionUser, login.Email)
		session.Save()
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusUnauthorized, nil)
	return
}

func (l *LoginController) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(SessionUser)
	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Abort()
	}
	c.Next()
}
