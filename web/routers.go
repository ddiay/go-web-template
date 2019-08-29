package web

import (
	"github.com/gin-gonic/gin"

	"github.com/ddiay/go-web-template/web/responses"
)

func initRouters(r *gin.Engine) {
	r.LoadHTMLGlob("assets/templates/*")

	// add routers here
	r.GET("/", responses.Index)

	r.Static("/static", "assets/static")
}
