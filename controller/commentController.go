package controller

import (
	"JeffMusic/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CreateComment  创建歌单
func CreateComment(c *gin.Context) {
	comment := new(models.Comment)
	json := make(map[string]interface{})
	c.BindJSON(&json)
	id, _ := c.Get("user_id")
	comment.UserId = id.(int)
	comment.SongId = json["song_id"].(int)
	comment.Comment = json["comment"].(string)
	comment.CreatedAt = json["created_at"].(time.Time)
	err := models.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

// GetComments 获取评论
func GetComments(c *gin.Context) {
	songIdString := c.Query("song_id")
	songId, _ := strconv.Atoi(songIdString)
	comments, err := models.GetComment(songId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, comments)
	}
}

// CreateReply  回复
func CreateReply(c *gin.Context) {
	reply := new(models.Reply)
	json := make(map[string]interface{})
	c.BindJSON(&json)
	id, _ := c.Get("user_id")
	reply.UserId = id.(int)
	reply.Reply = json["reply"].(string)
	reply.CreatedAt = json["created_at"].(time.Time)
	err := models.CreateReply(reply)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

// GetReplies 获取回复
func GetReplies(c *gin.Context) {
	commentIdString := c.Query("comment_id")
	commentId, _ := strconv.Atoi(commentIdString)
	replies, err := models.GetReply(commentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, replies)
	}
}
