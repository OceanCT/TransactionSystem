package Controller

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Service/BlogService"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddBlog(c *gin.Context) {
	request := Types.AddBlogRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response Types.AddBlogResponse
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = BlogService.AddBlog(id, request)
	c.JSON(http.StatusOK, response)
}
func ShowBlog(c *gin.Context) {
	request := Types.ShowBlogRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := BlogService.ShowBlog()
	c.JSON(http.StatusOK, response)
}
func Thumb(c *gin.Context) {
	request := Types.ThumbRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response Types.ThumbResponse
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = BlogService.Thumb(id, request)
	c.JSON(http.StatusOK, response)
}
func FindBlogThumb(c *gin.Context) {
	request := Types.FindBlogThumbRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response Types.FindBlogThumbResponse
	response = BlogService.FindBlogThumb(request)
	c.JSON(http.StatusOK, response)
}
