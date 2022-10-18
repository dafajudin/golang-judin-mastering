package routes

import (
	"go-hero/controllers"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//routing for hero
	e.GET("/hero", controllers.GetAllHero)
	e.POST("/hero", controllers.PostHero)
	e.PUT("/hero", controllers.UpdateHero)
	e.DELETE("/hero", controllers.DeleteHero)

	return e
}
