package Controller

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Service/VerifyCodeService"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddVerifyCode(c *gin.Context) {
	request := Types.AddVerifyCodeRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response = Types.AddVerifyCodeResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = VerifyCodeService.AddVerifyCode(id, request)
	c.JSON(http.StatusOK, response)
}
func ShowAllVerifyCode(c *gin.Context) {
	request := Types.ShowAllVerifyCodeRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response = Types.ShowAllVerifyCodeResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = VerifyCodeService.ShowVerifyCode(id)
	c.JSON(http.StatusOK, response)
}
func DeleteVerifyCode(c *gin.Context) {
	request := Types.DeleteVerifyCodeRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response = Types.DeleteVerifyCodeResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = VerifyCodeService.DeleteVerifyCode(id, request)
	c.JSON(http.StatusOK, response)
}
