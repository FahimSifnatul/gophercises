package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	uc := NewUseCase("redirect.yaml")
	NewHandler(e, uc)

	e.Logger.Fatal(e.Start(":8000"))
}
