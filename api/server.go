package api

import (
	"github.com/gobuffalo/pop"
	"github.com/labstack/echo/middleware"
	"github.com/uqichi/goffee-server/api/handler"
	"github.com/uqichi/goffee-server/usecases"

	"github.com/labstack/echo"
	"github.com/uqichi/goffee-server/repos"
)

func Serve() {
	e := echo.New()

	e.Debug = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := pop.Connect("development")
	if err != nil {
		panic(err)
	}

	repo := repos.NewCoffeeRepository(db)

	useCase := usecases.NewCoffeeUseCase(repo)

	h := handler.NewHandler(useCase)

	e.GET("/coffees/:id", h.GetCoffee)
	e.POST("/coffees", h.CreateCoffee)

	e.Logger.Fatal(e.Start(":1232"))
}
