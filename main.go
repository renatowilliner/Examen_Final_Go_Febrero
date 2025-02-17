package main

import (
	"examen/handler"
	"examen/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

	gestionHandler *handler.GestionHandler

)
func validateAuthHeader() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("x-caller-auth")

        if authHeader != "true" {
          
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Unauthorized: Invalid Token",
            })
            c.Abort() 
            return
        }

        c.Next()
    }
}
func main() {

	router = gin.Default()

	dependencies()

	mappingRoutes()

	router.Run(":8080")
}

func mappingRoutes() {
	router.Use(validateAuthHeader())
	groupGestion := router.Group("/gestion")
	groupGestion.POST("/calcularPromedioMensual", gestionHandler.CalcularPromedioMensual)
	groupGestion.POST("/calcularCostoConsumo", gestionHandler.CalcularCostoConsumo)
	groupGestion.GET("/generarProyeccionConsumo", gestionHandler.GenerarProyeccionConsumo)

}

func dependencies() {

	gestionService := service.NewGestionService()

	gestionHandler = handler.NewGestionHandler(gestionService)
	




}
