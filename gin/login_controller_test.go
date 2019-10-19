package gin_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xr1337/gin_bootstrap/gin"
)

func TestUserLogin(t *testing.T) {
	sut := gin.NewLoginController(nil)
	server := gin.NewServerDefault(nil, sut)

	t.Run("good user login", func(t *testing.T) {
		w := makeLogin(t, server.Router(), "john@gmail.com", "good")
		assertJSONHeaderResponse(t, w)
		assert.NotEmpty(t, w.Header().Get("Set-Cookie"))
	})

	t.Run("bad user login", func(t *testing.T) {
		testAccount := []gin.Login{
			gin.Login{Email: "john@gmail.com", Password: "bad"},
			gin.Login{Email: "", Password: ""},
		}
		for _, i := range testAccount {
			w := makeLogin(t, server.Router(), i.Email, i.Password)
			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.Empty(t, w.Header().Get("Set-Cookie"))
		}
	})
}

func TestAccessAboutMe(t *testing.T) {
	sut := gin.NewLoginController(nil)
	server := gin.NewServerDefault(nil, sut)

	t.Run("Can access Me after login", func(t *testing.T) {
		w := makeLogin(t, server.Router(), "john@gmail.com", "good")
		assertJSONHeaderResponse(t, w)
		cookie := w.Header().Get("Set-Cookie")

		assert.NotEmpty(t, cookie)
		w = performRequest(t, server.Router(), "GET", "/user/me", nil, cookie)
		assertJSONHeaderResponse(t, w)
	})

	t.Run("Can unauthorize Me", func(t *testing.T) {
		w := performRequest(t, server.Router(), "GET", "/user/me", nil, "bad cookie")
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func makeLogin(t *testing.T, server http.Handler, email string, password string) *httptest.ResponseRecorder {
	t.Helper()
	login := gin.Login{Email: email, Password: password}
	return performRequest(t, server, "POST", "/user/login", toJSONReader(login), "")
}
