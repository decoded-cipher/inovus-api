package main

import (
	"strconv"

	"github.com/decoded-cipher/inovus-api/internal"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getAllUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	users := []User{
		{ID: uuid.Must(uuid.NewV4()).String(), Name: "Alice"},
		{ID: uuid.Must(uuid.NewV4()).String(), Name: "Bob"},
		{ID: uuid.Must(uuid.NewV4()).String(), Name: "Charlie"},
	}

	response := internal.Paginate(users, page, pageSize)
	return c.JSON(200, response)
}
