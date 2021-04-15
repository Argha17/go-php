package httphandler

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	return router
}
