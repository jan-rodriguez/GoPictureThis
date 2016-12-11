package utils

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func SendErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
