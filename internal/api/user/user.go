package user

import (
	"fmt"
	"net/http"
	"time"
	"todo/helpers"
	"todo/internal/config"
	"todo/internal/models"
	"todo/repository"
	"todo/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var req request.Login
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user, err := repository.Get().User().Login(req.UserName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "bir hatavar "))
	}

	passwordControl := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passwordControl != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Şifre doğrulanma"))
	}

	claims := &config.JwtCustom{
		User: *&user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := Token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helpers.Response(err, "token oluşturma hatası"))
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t, "user": user})
}

func Register(c echo.Context) error {
	var req request.Register

	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	// duplicate user control
	count := repository.Get().User().DuplicateUserName(req.UserName)
	if count > 0 {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "bu kullanıcı adı kullanılmakta"))
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)
	user := models.User{
		Name:     req.Name,
		Surname:  req.Surname,
		UserName: req.UserName,
		Password: string(password),
	}

	err := repository.Get().User().Insert(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Kayıt işlemi başarısız"))
	}

	return c.JSON(http.StatusOK, helpers.Response(user, "Kayıt başarılı"))
}
