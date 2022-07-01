package main

import (
	"fmt"
	"os"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/victorhsb/go-mocker/http"
	"github.com/victorhsb/go-mocker/service"
	"github.com/victorhsb/go-mocker/storage"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	stg := storage.NewStorage(path.Join(wd, "example"))
	svc := service.NewService(stg)

	routes, err := svc.GetRoutes()
	if err != nil {
		panic(err)
	}

	ec := echo.New()
	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())

	for _, r := range routes {
		ec.Add(r.Method, r.Path, http.Handler(r))
		fmt.Printf("registered route %s %s\n", r.Method, r.Path)
	}

	ec.Logger.Fatal(ec.Start(":8080"))
}
