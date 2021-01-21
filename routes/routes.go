package routes

import (
	"net/http"
	"github.com/labstack/echo"
	"golang-crud-test/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "WHAT? HOW?")
	})

	e.GET("/user", controllers.FetchAllUser)
	e.POST("/user/", controllers.StoreUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	return e
}
