package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.POST("/users", createUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println(u)

	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/echodb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return err
	}

	db.Create(&u)
	defer db.Close()

	return c.JSON(http.StatusOK, u)
}
