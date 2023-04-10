package router

import (
	"github.com/finallly/srp_protocol/web/internal/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/static/*")
	handlers.InitHandlers(router)
	return router
}
