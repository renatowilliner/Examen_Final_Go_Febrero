package handler

import (
	"examen/model"
	"examen/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GestionHandler struct {
	gestionService service.GestionInterface
}

func NewGestionHandler(gestionService service.GestionInterface) *GestionHandler {
	return &GestionHandler{gestionService: gestionService}
}


func (handler *GestionHandler) CalcularPromedioMensual(c *gin.Context) {
	var consumo model.ConsumoElectrico

	if handler.gestionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servicio no configurado"})
		return
	}

	err := c.ShouldBindJSON(&consumo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, resultado := handler.gestionService.CalcularPromedioMensual(&consumo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"totalConsumo": resultado})
	}
}


func (handler *GestionHandler) CalcularCostoConsumo(c *gin.Context) {
	var consumoMensual model.ConsumoElectricoMensual

	if handler.gestionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servicio no configurado"})
		return
	}

	err := c.ShouldBindJSON(&consumoMensual)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, resultado := handler.gestionService.CalcularCostoConsumo(&consumoMensual)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"costoTotal": resultado})
	}
}


func (handler *GestionHandler) GenerarProyeccionConsumo(c *gin.Context) {
	consumoMensual, err1 := strconv.ParseFloat(c.Query("consumoMensual"), 64)
	tasaAumentoAnual, err2 := strconv.ParseFloat(c.Query("tasaAumentoAnual"), 64)
	anios, err3 := strconv.Atoi(c.Query("anios"))

	if handler.gestionService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servicio no configurado"})
		return
	}

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	err, resultado := handler.gestionService.GenerarProyeccionConsumo(consumoMensual, tasaAumentoAnual, anios)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"proyeccionConsumo": resultado})
	}
}
