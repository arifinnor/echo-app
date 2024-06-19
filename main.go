package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "User 1"},
	{ID: 2, Name: "User 2"},
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", getAllUsers)
	e.GET("/user/:id", getUserByID)

	e.Logger.Fatal(e.Start(":1323"))
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, u := range users {
		if u.ID == id {
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.String(http.StatusNotFound, "User not found")
}
