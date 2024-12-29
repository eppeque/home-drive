package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
