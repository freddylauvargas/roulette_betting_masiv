package roulette

import "github.com/labstack/echo"

func BettingRouter(e *echo.Echo) {

	r := e.Group("/Betting")
	r.POST("/NewRoulette", NewRoulette)
	r.POST("/OpenRoulette", OpenRoulette)
	r.POST("/BetNumberRoulette", BetNumberRoulette)
	r.POST("/BetClosingRoulette", BetClosingRoulette)
	r.POST("/ListRoulette", ListRoulette)

}