package roulette

type requestNewRoulette struct {
	Function string `json:"function"`
}
type responseNewRoulette struct {
	Status     int    `json:"status"`
	IDRoulette string `json:"id_roulette"`
	CreatedAt  string `json:"created_at"`
}