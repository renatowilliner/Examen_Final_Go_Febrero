package model

type ConsumoElectricoMensual struct {
	ConsumoMensual float64 `json:"consumoMensual"`
	CostoPorKWh    float64 `json:"costoPorKWh"`
}
