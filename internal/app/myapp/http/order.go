package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Order(c *gin.Context) {
	c.JSON(http.StatusOK,svc.Orders())
}
