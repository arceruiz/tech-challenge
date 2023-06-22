package rest

import (
	"client/internal/canonical"
	"client/internal/middlewares"
	"client/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type rest struct {
	svc service.Service
}

func New() *rest {
	return &rest{
		svc: service.New(),
	}
}

func (r *rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	router.GET("/register", r.Register)
	router.GET("/login", r.Login)
	router.GET("/bypass", echo.MethodNotAllowedHandler)

	return router.Start(":3001")
}

func (r *rest) Register(c echo.Context) error {
	var userRequest User

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(400, "dados incorretos")
	}

	user, err := r.svc.Register(canonical.User{
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

func (r *rest) Login(c echo.Context) error {
	var userLogin User

	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(400, "Login ou senha incorretos")
	}

	token, err := r.svc.Login(canonical.User{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, Response{
		Token: token,
	})
}
