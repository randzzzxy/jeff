package controller

import (
	"JeffMusic/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreatePlayList 创建歌单
func CreatePlayList(c *gin.Context) {
	playList := new(models.PlayList)
	id, _ := c.Get("user_id")
	playList.UserId = id.(int)
	playList.Name = c.PostForm("play_list_name")
	err := models.CreatePlayList(playList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, playList)
	}
}

// GetPlayLists 获取歌单
func GetPlayLists(c *gin.Context) {
	id, _ := c.Get("user_id")
	playLists, err := models.GetPlayList(id.(int))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, playLists)
	}
}

// CollectSongToPlayList 收藏歌曲
func CollectSongToPlayList(c *gin.Context) {
	var collect models.Collect
	c.BindJSON(&collect)
	err := models.CreateCollect(&collect)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, collect)
	}
}

// GetPlayList 获取歌单歌曲
func GetPlayList(c *gin.Context) {
	idString := c.Query("play_list_id")
	id, _ := strconv.Atoi(idString)
	list, err := models.GetPlayListSongs(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, list)
	}
}
