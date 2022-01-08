package todo

import (
	"net/http"
	"todo/helpers"
	"todo/internal/models"
	"todo/repository"
	"todo/request"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	user := helpers.Auth(&c)

	todos, err := repository.Get().Todo().List(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, ""))
	}
	return c.JSON(http.StatusOK, helpers.Response(todos, ""))
}

func Insert(c echo.Context) error {
	var req request.TodoInsert
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user := helpers.Auth(&c)

	todo := models.Todo{
		Userfk: user.ID,
		Task:   req.Task,
	}

	err := repository.Get().Todo().Insert(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kaydetme başarısız"))
	}

	return c.JSON(http.StatusOK, helpers.Response(todo, "Başarılı"))
}
