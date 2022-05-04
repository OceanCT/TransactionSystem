package Types

import (
	"HomeworkForDB/Types/Const/ErrNo"
	"HomeworkForDB/Types/Const/ThumbStatus"
	"HomeworkForDB/Types/Const/TransactionStatus"
)

type RegisterUserRequest struct {
	User       User   `json:"user" binding:"required"`
	VerifyCode string `json:"verifyCode" binding:"required"`
}
type RegisterUserResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message  string   `json:"message" binding:"required"`
		User     User     `json:"user" binding:"required"`
		Userinfo UserInfo `json:"userinfo" `
	}
}
type LogoutRequest struct{}
type LogoutResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type WhoAmIRequest struct{}
type WhoAmIResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message  string   `json:"message" binding:"required"`
		User     User     `json:"user" binding:"required"`
		Userinfo UserInfo `json:"userinfo" `
	}
}
type ChangeUserRequest struct {
	User User `json:"user" binding:"required"`
}
type ChangeUserResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type ChangeUserInfoRequest struct {
	UserInfo UserInfo `json:"userInfo" binding:"required"`
}
type ChangeUserInfoResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type AddVerifyCodeRequest struct {
	VerifyCode VerifyCode `json:"verifyCode" binding:"required"`
}
type AddVerifyCodeResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type ShowAllVerifyCodeRequest struct{}
type ShowAllVerifyCodeResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message     string       `json:"message" binding:"required"`
		VerifyCodes []VerifyCode `json:"verifyCodes" binding:"required"`
	}
}
type DeleteVerifyCodeRequest struct {
	VerifyCode string `json:"verifyCode" binding:"required"`
}
type DeleteVerifyCodeResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type AddBlogRequest struct {
	Blog Blog `json:"blog" binding:"required"`
}
type AddBlogResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type ShowBlogRequest struct{}
type ShowBlogResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
		Blogs   []Blog `json:"blogs" binding:"required"`
	}
}
type ShowBillRequest struct{}
type ShowBillResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
		Bills   []Bill `json:"bills" binding:"required"`
	}
}
type PayBillRequest struct {
	BillID int64 `json:"billID" binding:"required"`
}
type PayBillResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type ShowTransactionRequest struct{}
type ShowTransactionResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message      string        `json:"message" binding:"required"`
		Transactions []Transaction `json:"transactions" binding:"required"`
	}
}
type SetTransactionStatusRequest struct {
	TransactionID     int64                               `json:"transactionId" binding:"required"`
	TransactionStatus TransactionStatus.TransactionStatus `json:"transactionStatus" binding:"required"`
}
type SetTransactionStatusResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type CreateTransactionRequest struct {
	TransactionInfo string `json:"transactionInfo" binding:"required"`
	To              int64  `json:"to" binding:"required"`
	Bill            Bill   `json:"bill" binding:"required"`
}
type CreateTransactionResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type ThumbRequest struct {
	BlogID int64                   `json:"blogID" binding:"required"`
	Status ThumbStatus.ThumbStatus `json:"status" binding:"required"`
}
type ThumbResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message string `json:"message" binding:"required"`
	}
}
type FindBlogThumbRequest struct{ BlogID int64 }
type FindBlogThumbResponse struct {
	Code ErrNo.ErrNo `json:"code" binding:"required"`
	Data struct {
		Message   string `json:"message" binding:"required"`
		ThumbSum1 int    `json:"thumbUp" binding:"required"`
		ThumbSum2 int    `json:"thumbDown" binding:"required"`
	}
}
