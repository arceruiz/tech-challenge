package port

import (
	"client/internal/canonical"

	"client/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type UserPort struct {
	service service.UserService
}

func NewUserPort(userService service.UserService) *UserPort {
	return &UserPort{
		service: userService,
	}
}

func (u *UserPort) Register(g *echo.Group) {
	g.GET("/register", u.register)
	g.GET("/login", u.login)
	g.GET("/bypass", echo.MethodNotAllowedHandler)
}

func (u *UserPort) register(c echo.Context) error {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(400, "dados incorretos")
	}

	user, err := u.service.Register(canonical.User{
		Document:  userRequest.Document,
		Email:     userRequest.Email,
		Name:      userRequest.Name,
		Password:  userRequest.Password,
		CreatedAt: time.Now().String(),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, "implementar um map pra erros, algo assim")
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserPort) login(c echo.Context) error {
	var userLogin UserRequest

	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(400, "Login ou senha incorretos")
	}

	token, err := u.service.Login(canonical.User{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
