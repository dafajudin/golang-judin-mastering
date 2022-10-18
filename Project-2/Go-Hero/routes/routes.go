package routes

import (
	"go-hero/controllers"
	"net/http"
	"github.com/labstack/echo/v4"
	"go-hero/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//routing for hero
	e.GET("/hero", controllers.GetAllHero,  middleware.IsAuthenticated)
	e.POST("/hero", controllers.PostHero, middleware.IsAuthenticated)
	e.PUT("/hero", controllers.UpdateHero, middleware.IsAuthenticated)
	e.DELETE("/hero", controllers.DeleteHero, middleware.IsAuthenticated)

	//routing for login
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLoginUser)

	return e
}
