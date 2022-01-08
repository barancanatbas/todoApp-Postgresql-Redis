package home

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	fmt.Println("selam ")
	return nil
}
