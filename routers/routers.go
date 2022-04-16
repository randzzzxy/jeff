package routers

import (
	"JeffMusic/controller"
	"JeffMusic/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 用户
	userGroup := r.Group("user")
	{
		// 注册
		userGroup.POST("/register", controller.RegisterNewAccount)
		// 登陆
		userGroup.POST("/login", controller.Login)
	}
	musicGroup := r.Group("music")
	musicGroup.Use(controller.ValidateTokenHandler)
	{
		// 上传歌曲
		musicGroup.POST("/upload")
	}
	return r
}
