package Controller

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Service/TransactionService"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateTransaction(c *gin.Context) {
	var request Types.CreateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := Types.CreateTransactionResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = TransactionService.CreateTransaction(id, request)
	c.JSON(http.StatusOK, response)
}
func SetTransactionStatus(c *gin.Context) {
	var request Types.SetTransactionStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := Types.SetTransactionStatusResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = TransactionService.SetTransactionStatus(id, request)
	c.JSON(http.StatusOK, response)
}
func ShowTransaction(c *gin.Context) {
	var request Types.ShowTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := Types.ShowTransactionResponse{}
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = TransactionService.ShowTransactions(id)
	c.JSON(http.StatusOK, response)
}
