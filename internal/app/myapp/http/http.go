package http

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"orp/internal/app/myapp/service"
)

var (
	svc *service.Service
)

func Init(c string, s *service.Service) *gin.Engine {
	svc = s
	return innerRouter()
}
func innerRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	//gin.DisableConsoleColor()
	gin.DefaultWriter = ioutil.Discard //把日志丢弃掉
	router = gin.Default()

	//r.Use(logger.Logger(conf.Cfg.Log))

	//开启gzip压缩
	//router.Use(gzip.Gzip(gzip.BestCompression))
	router.GET("/", func(ctx *gin.Context) {
		//ctx.String(http.StatusOK, "abs")
		var du float64 = 1.00
		ctx.JSON(http.StatusOK, gin.H{"aa": du})
	})
	router.GET("/order",Order)
	return
}
