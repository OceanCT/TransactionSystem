package Controller

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Service/BillService"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PayBill(c *gin.Context) {
	request := Types.PayBillRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response Types.PayBillResponse
	response = BillService.PayBill(request)
	c.JSON(http.StatusOK, response)
}
func ShowBill(c *gin.Context) {
	request := Types.ShowBillRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var response Types.ShowBillResponse
	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.ErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	var id int64
	id, err = strconv.ParseInt(cookie.Value, 10, 64)
	response = BillService.ShowBill(id)
	c.JSON(http.StatusOK, response)
}
