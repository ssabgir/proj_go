package main

import (
	"awesomeProject/accounts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.PATCH("/account/change/amount", accountsHandler.ChangeAccountBalance)
	e.PATCH("/account/change/name", accountsHandler.ChangeAccountName)
	e.DELETE("/account/delete", accountsHandler.DeleteAccount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
