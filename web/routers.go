package web

import (
	"github.com/gin-gonic/gin"

	"github.com/ddiay/go-web-template/web/responses"
)

func initRouters(r *gin.Engine) {
	r.LoadHTMLGlob("assets/templates/*")

	// add routers here
	r.GET("/", responses.Index)

	// js目录
	r.Static("/js", "assets/js")

	// css目录
	r.Static("/css", "assets/css")

	// images目录
	r.Static("/images", "assets/images")

	// 静态文件目录
	r.Static("/static", "assets/static")

	// files目录
	r.Static("/files", "assets/files")
}
