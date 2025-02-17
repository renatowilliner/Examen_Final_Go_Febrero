package service

import (
	"errors"
	"examen/model"
)

type GestionInterface interface {
	CalcularPromedioMensual(*model.ConsumoElectrico) (error, float64)
	CalcularCostoConsumo(*model.ConsumoElectricoMensual) (error, float64)
	GenerarProyeccionConsumo(float64, float64, int) (error, []model.ProyeccionConsumo)
}

type GestionService struct{}

func NewGestionService() GestionInterface {
	return &GestionService{}
}

func (service *GestionService) CalcularPromedioMensual(consumoElectrico *model.ConsumoElectrico) (error, float64) {
	if consumoElectrico == nil || len(consumoElectrico.Consumos) == 0 {
		return errors.New("los datos de consumo eléctrico son nulos"), 0
	}

	var suma float64
	for _, consumo := range consumoElectrico.Consumos {
		if consumo < 0 {
			return errors.New("los valores de consumo no pueden ser negativos"), 0
		}
		suma += consumo
	}

	promedio := suma / float64(len(consumoElectrico.Consumos))
	return nil, promedio
}

func (service *GestionService) CalcularCostoConsumo(consumoElectricoMensual *model.ConsumoElectricoMensual) (error, float64) {
	if consumoElectricoMensual == nil {
		return errors.New("los datos de consumo eléctrico mensual son nulos"), 0
	}
	if consumoElectricoMensual.ConsumoMensual < 0 {
		return errors.New("el consumo mensual no puede ser negativo"), 0
	}
	if consumoElectricoMensual.CostoPorKWh < 0 {
		return errors.New("el costo por kWh no puede ser negativo"), 0
	}

	costoTotal := consumoElectricoMensual.ConsumoMensual * consumoElectricoMensual.CostoPorKWh
	return nil, costoTotal
}

func (service *GestionService) GenerarProyeccionConsumo(consumoMensual float64, tasaAumentoAnual float64, anios int) (error, []model.ProyeccionConsumo) {
	if consumoMensual < 0 {
		return errors.New("el consumo mensual no puede ser negativo"), nil
	}
	if tasaAumentoAnual < 0 || tasaAumentoAnual > 100 {
		return errors.New("la tasa de aumento debe estar entre 0 y 100"), nil
	}
	if anios <= 0 {
		return errors.New("el número de años debe ser mayor a 0"), nil
	}

	var proyeccion []model.ProyeccionConsumo
	consumoActual := consumoMensual

	for i := 1; i <= anios; i++ {
		consumoActual = consumoActual * (1 + tasaAumentoAnual/100)
		proyeccion = append(proyeccion, model.ProyeccionConsumo{
			Anio:    i,
			Consumo: consumoActual,
		})
	}

	return nil, proyeccion
}
