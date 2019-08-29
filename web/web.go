package web

import (
	"io"
	"os"

	"github.com/ddiay/go-log"

	"github.com/ddiay/go-web-template/web/common/config"
	"github.com/ddiay/go-web-template/web/common/logger"

	"github.com/gin-gonic/gin"
)

func StartWeb() bool {
	// 初始化Web日志对象
	logger.WebLogger = log.NewLogger().AppendWriters(log.NewFileWriter().Filename("logs/web/web.log"))
	defer logger.WebLogger.Close()

	// 载入程序配置
	err := config.LoadConfig("web.cfg")
	if err != nil {
		logger.Warn(err.Error())
		// return
	}

	if !config.WebCfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	ginlogWriter := log.NewFileWriter().Filename("logs/web/gin.log")
	gin.DefaultWriter = io.MultiWriter(ginlogWriter, os.Stdout)

	r := gin.Default()

	initRouters(r)

	err = r.Run(":" + config.WebCfg.Port)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
		return false
	}

	return true
}
