package todo

import (
	"net/http"
	"strconv"
	"todo/cache"
	"todo/helpers"
	"todo/internal/models"
	"todo/repository"
	"todo/request"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	user := helpers.Auth(&c)

	var todos []models.Todo
	redisKey := "todos-" + strconv.Itoa(int(user.ID))
	if cache.GetFromCache(redisKey, &todos) {
		return c.JSON(http.StatusOK, helpers.Response(todos, ""))
	}

	todos, err := repository.Get().Todo().List(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, ""))
	}

	cache.SetFromCache(redisKey, todos, 3600)
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

	// yeni todo eklendiği zaman eski verileri temizlemek gerek
	redisKey := "todos-" + strconv.Itoa(int(user.ID))
	cache.DelFromCache(redisKey)

	return c.JSON(http.StatusOK, helpers.Response(todo, "Başarılı"))
}

func Completed(c echo.Context) error {
	var req request.Completed

	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user := helpers.Auth(&c)
	// check todo
	todo, err := repository.Get().Todo().GetTodo(req.ID, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Todo bulunamadı"))
	}

	todo.Completed = true

	err = repository.Get().Todo().Update(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Todo güncellenemedi"))
	}

	// veri güncellemesi yapıldığı zaman cacheden veri sil
	redisKey := "todos-" + strconv.Itoa(int(user.ID))
	cache.DelFromCache(redisKey)

	return c.JSON(http.StatusOK, helpers.Response(nil, "Todo tamamlandı"))
}

func Delete(c echo.Context) error {
	var req request.Completed

	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user := helpers.Auth(&c)
	// Delete todo
	err := repository.Get().Todo().Delete(req.ID, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Todo Silinemedi"))
	}

	// eski veriyi sildiği zaman cacheden veri sil
	redisKey := "todos-" + strconv.Itoa(int(user.ID))
	cache.DelFromCache(redisKey)

	return c.JSON(http.StatusOK, helpers.Response(nil, "Todo silindi"))
}
