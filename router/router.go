package router

import (
	"github.com/gin-gonic/gin"
	"github.com/inontz/kepler-go/handler"
)

func Routes(r *gin.Engine) {

	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.GET("/", handler.Index)

}
