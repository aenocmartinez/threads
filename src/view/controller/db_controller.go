package controller

import (
	"net/http"
	"threads/src/infraestructure/database"

	"github.com/gin-gonic/gin"
)

func CheckDBConnection(c *gin.Context) {
	db := database.GetDB()

	if err := db.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "No se pudo conectar a la base de datos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Conexión a la base de datos establecida correctamente",
	})
}
