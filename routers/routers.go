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
	//文件
	fileGroup := r.Group("file")
	{
		//音乐文件
		fileGroup.GET("/song", controller.GetSongResource)
		//图片
		fileGroup.GET("/cover", controller.GetCoverResource)
		//歌词文件
		fileGroup.GET("/lyrics", controller.GetLyricsResource)
	}
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
		//获取歌曲
		musicGroup.GET("/all", controller.GetSongs)
		// 创建歌单
		musicGroup.POST("/playlist", controller.CreatePlayList)
		// 获取歌单
		musicGroup.GET("/playlist", controller.GetPlayLists)
		// 收藏歌曲
		musicGroup.POST("/collect", controller.CollectSongToPlayList)
		// 获取歌单歌曲
		musicGroup.GET("/collect", controller.GetPlayList)
		// 上传歌曲
		musicGroup.POST("/upload", controller.UploadMusic)
	}
	commentGroup := r.Group("comment")
	commentGroup.Use(controller.ValidateTokenHandler)
	{
		// 评论
		commentGroup.POST("/comment", controller.CreateComment)
		// 获取评论
		commentGroup.GET("/comment", controller.GetComments)
		// 回复
		commentGroup.POST("/reply", controller.CreateReply)
		// 获取回复
		commentGroup.GET("/reply", controller.GetReplies)
	}
	return r
}
