package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masiv/internal/configuration"
	"github.com/masiv/roulette"
)

func main(){
	fmt.Println("Version 1.0.2")
	c := configuration.FromFile()
	// Definicion de Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// creacion de rutas
	roulette.BettingRouter(e)

	// levantar el servicio
	fmt.Println("Start server Version 1.0.2")
	e.Logger.Fatal(e.Start(":" + c.AppPort))
}