package router

import (
	t "todo/internal/api/todo"
	user "todo/internal/api/user"
	"todo/internal/config"
	_middleware "todo/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	// route
	e.POST("/login", user.Login)
	e.POST("/register", user.Register)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_middleware.Auth)

	todo := admin.Group("")
	todo.GET("/todos", t.Index)
	todo.POST("/todo", t.Insert)
	todo.GET("/todo/completed", t.Completed)
	todo.DELETE("/todo", t.Delete)
}
