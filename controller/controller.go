package controller

import (
	"JeffMusic/models"
	"JeffMusic/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func RegisterNewAccount(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var user models.User
	c.BindJSON(&user)
	// 2. 存入数据库
	err := models.CreateNewAccount(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user, err := models.GetUserInfo(u.UnionID)
	if err != nil || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := utils.CreateToken(user.UnionID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func ValidateTokenHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusForbidden, "The authorization token is abnormal.")
		c.Abort()
		return
	}
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}
	c.Set("user_id", claims.UserId)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}

func UploadMusic(c *gin.Context) {
	id, _ := c.Get("user_id")
	coverFile, coverErr := c.FormFile("cover")
	lyricsFile, lyricsErr := c.FormFile("lyrics")
	songFile, songErr := c.FormFile("song")
	if coverErr != nil {
		c.String(500, "上传封面出错")
		c.Abort()
		return
	}
	if songErr != nil {
		c.String(500, "上传歌曲出错")
		c.Abort()
		return
	}
	song := new(models.Song)
	song.AuthorId = id.(int)
	song.SongUrl = songFile.Filename
	song.CoverUrl = coverFile.Filename
	if lyricsErr == nil {
		if err := c.SaveUploadedFile(lyricsFile, "./lyrics/"+lyricsFile.Filename); err != nil {
			fmt.Println(err)
		}
		song.LyricsUrl = lyricsFile.Filename
	}
	models.UploadSong(song)
	// c.JSON(200, gin.H{"message": file.Header.Context})
	if err := c.SaveUploadedFile(songFile, "./songs/"+songFile.Filename); err != nil {
		fmt.Println(err)
	}
	if err := c.SaveUploadedFile(coverFile, "./covers/"+coverFile.Filename); err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, songFile.Filename)
}

func GetSongResource(c *gin.Context) {
	c.File("./songs/" + c.Query("name"))
}

func GetCoverResource(c *gin.Context) {
	c.File("./covers/" + c.Query("name"))
}

func GetLyricsResource(c *gin.Context) {
	c.File("./lyrics/" + c.Query("name"))
}
