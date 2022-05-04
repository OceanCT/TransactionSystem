package Router

import (
	"HomeworkForDB/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	// auth
	g.POST("/auth/login", Controller.Login)
	g.POST("/auth/logout", Controller.Logout)
	// user
	g.POST("/user/register", Controller.RegisterUser)
	g.POST("/user/changeUser", Controller.ChangeUser)
	g.POST("/user/changeUserInfo", Controller.ChangeUserInfo)
	g.POST("/user/whoAmI", Controller.WhoAmI)
	// verifyCode
	g.POST("/verifyCode/addVerifyCode", Controller.AddVerifyCode)
	g.POST("/verifyCode/deleteVerifyCode", Controller.DeleteVerifyCode)
	g.POST("/verifyCode/showVerifyCode", Controller.ShowAllVerifyCode)
	// blog
	g.POST("/blog/addBlog", Controller.AddBlog)
	g.POST("/blog/showBlog", Controller.ShowBlog)
	g.POST("/blog/findBlogThumb", Controller.FindBlogThumb)
	// bill
	g.POST("/bill/payBill", Controller.PayBill)
	g.POST("/bill/showBill", Controller.ShowBill)
	// transaction
	g.POST("/transaction/createTransaction", Controller.CreateTransaction)
	g.POST("/transaction/showTransaction", Controller.ShowTransaction)
	g.POST("/transaction/setTransactionStatus", Controller.SetTransactionStatus)
	// thumb
	g.POST("/thumb/thumb", Controller.Thumb)
}
