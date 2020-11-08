package roulette

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/masiv/internal/logger"
	"net/http"
)

func NewRoulette(c echo.Context) error {
	fmt.Println("Crear Ruleta")
	requestNewRoulette := requestNewRoulette{}
	responseNewRoulette := responseNewRoulette{}
	err := c.Bind(&requestNewRoulette)
	if err != nil {
		logger.Error.Println("The struct of the object no valid : ", err)
		return c.JSON(http.StatusAccepted, responseNewRoulette)
	}

	responseNewRoulette.Status = 200
	responseNewRoulette.IDRoulette = "AC051938-53C3-43EA-8636-59E485EC1007"
	responseNewRoulette.CreatedAt = "2020-11-08T16:27:01Z"

	return c.JSON(http.StatusOK, responseNewRoulette)
}
func OpenRoulette(c echo.Context) error {
	fmt.Println("Abrir Ruleta")
	return nil
}
func BetNumberRoulette(c echo.Context) error {
	fmt.Println("Ruleta - Apuesta a un Número")
	return nil
}
func BetClosingRoulette(c echo.Context) error {
	fmt.Println("Cerrar Ruleta")
	return nil
}
func ListRoulette(c echo.Context) error {
	fmt.Println("Listar Ruletas")
	return nil
}
