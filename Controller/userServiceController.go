package Controller

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Service/UserService"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	request := Types.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := UserService.Login(request)
	c.SetCookie("camp-session", strconv.FormatInt(response.Data.User.UserID, 10), 0, "/", "", false, true)
	c.JSON(http.StatusOK, response)
}
func Logout(c *gin.Context) {
	response := &Types.LogoutResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err == nil {
		c.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		response.Code = ErrNo.OK
	} else {
		response.Code = ErrNo.LoginRequired
	}
	c.JSON(http.StatusOK, response)
}
func WhoAmI(c *gin.Context) {
	request := Types.WhoAmIRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := Types.WhoAmIResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response = UserService.WhoAmI(id)
	c.JSON(http.StatusOK, response)
}
func RegisterUser(c *gin.Context) {
	request := Types.RegisterUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := UserService.RegisterUser(request)
	c.JSON(http.StatusOK, response)
}
func ChangeUser(c *gin.Context) {
	request := Types.ChangeUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := Types.ChangeUserResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response = UserService.UpdateUser(id, request)
	c.JSON(http.StatusOK, response)
}
func ChangeUserInfo(c *gin.Context) {
	request := Types.ChangeUserInfoRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response := Types.ChangeUserInfoResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	response = UserService.UpdateUserInfo(id, request)
	c.JSON(http.StatusOK, response)
}
