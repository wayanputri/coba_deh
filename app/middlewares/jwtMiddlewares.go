package middlewares

import (
	"alta/immersive-dashboard-api/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var appConfig = config.ReadEnv()

func JWTMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(appConfig.JWT_ACCESS_TOKEN),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int) (string, error){
	claims := jwt.MapClaims{}
	claims["authorized"] =true
	claims["userId"] = userId
	claims["exp"]= time.Now().Add(time.Hour*1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(appConfig.JWT_ACCESS_TOKEN))
}

func ExtracTokenUserId(e echo.Context) int{
	user := e.Get("user").(*jwt.Token)
	if user.Valid{
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}