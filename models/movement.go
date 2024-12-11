package models

type Velocity struct {
	Vx    float64 `json:"vx"`
	Vy    float64 `json:"vy"`
	Omega float64 `json:"omega"`
}

type AgvPosition struct {
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"`
	PositionInitialized bool    `json:"positionInitialized"`
	MapDescription      string  `json:"mapDescription"`
	LocalizationScore   float64 `json:"localizationScore"`
	DeviationRange      float64 `json:"deviationRange"`
}
