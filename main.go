package main

import (
	user "github.com/dovics/catapult/view/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v8"
)

var engine *xorm.Engine

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	/*e.Use(middleware.JWT(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "query:token",
	}))*/
	initDB()
	initRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func initDB() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(10.0.0.251:3306)/community?charset=utf8")
	if err != nil {
		return
	}
}

// init router
func initRouter(e *echo.Echo) {
	config := validator.Config{}
	validator := validator.New(&config)

	userController := user.New(engine, validator)

	e.POST("/user/register", userController.Register)
}
