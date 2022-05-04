package Types

import (
	"HomeworkForDB/Types/Const/BillStatus"
	"HomeworkForDB/Types/Const/TransactionStatus"
	"HomeworkForDB/Types/Const/UserType"
)

type User struct {
	UserID   int64             `json:"userID"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	UserType UserType.UserType `json:"userType"`
	Money    float64           `json:"money" `
}

type UserInfo struct {
	UserID    int64  `json:"userID"`
	Birthday  string `json:"birthday"`
	DeathDate string `json:"deathDate"`
	UserInfo  string `json:"userInfo"`
}

type Blog struct {
	ID        int64  `json:"blogID" `
	BlogInfo  string `json:"blogInfo" binding:"required"`
	CreatorID int64  `json:"creatorID"`
}

type Thumb struct {
	BlogID int64 `json:"blogID" binding:"required"`
	UserID int64 `json:"userID" binding:"required"`
	Thumb  int   `json:"thumb" binding:"required"`
}

type Bill struct {
	ID        int64                 `json:"id" `
	UserID    int64                 `json:"userID" binding:"required"`
	BillMoney float64               `json:"billMoney" binding:"required"`
	Status    BillStatus.BillStatus `json:"status" `
}

type Transaction struct {
	ID                int64                               `json:"id" binding:"required"`
	TransactionInfo   string                              `json:"transactionInfo" binding:"required"`
	From              string                              `json:"from" binding:"required"`
	To                string                              `json:"to" binding:"required"`
	TransactionStatus TransactionStatus.TransactionStatus `json:"transactionStatus" binding:"required"`
}

type VerifyCode struct {
	VerifyCode string            `json:"verifyCode" binding:"required"`
	UserType   UserType.UserType `json:"userType" binding:"required"`
}
