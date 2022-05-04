package VerifyCodeService

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Dao/UserDao"
	"HomeworkForDB/Dao/VerifyCodeDao"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"HomeworkForDB/Types/Const/UserType"
	"database/sql"
)

var db *sql.DB
var DBErr error

func init() {
	db, DBErr = DBAccessor.MysqlInit()
}

func AddVerifyCode(userid int64, request Types.AddVerifyCodeRequest) Types.AddVerifyCodeResponse {
	var res Types.AddVerifyCodeResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	tx, err := db.Begin()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	var user = Types.User{UserID: userid}
	res.Code = UserDao.FindUserByID(tx, &user)
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	if user.UserType != UserType.Admin {
		res.Code = ErrNo.PermDenied
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	res.Code = VerifyCodeDao.InsertVerifyCode(tx, request.VerifyCode)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func ShowVerifyCode(userid int64) Types.ShowAllVerifyCodeResponse {
	var res Types.ShowAllVerifyCodeResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var user = Types.User{UserID: userid}
	tx, err := db.Begin()
	res.Code = UserDao.FindUserByID(tx, &user)
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	if res.Code == ErrNo.OK {
		if user.UserType != UserType.Admin {
			res.Code = ErrNo.PermDenied
		} else {
			res.Code = VerifyCodeDao.FindAllVerifyCode(tx, &res.Data.VerifyCodes)
		}
	}
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func DeleteVerifyCode(userid int64, request Types.DeleteVerifyCodeRequest) Types.DeleteVerifyCodeResponse {
	var res Types.DeleteVerifyCodeResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var user = Types.User{UserID: userid}
	tx, err := db.Begin()
	res.Code = UserDao.FindUserByID(tx, &user)
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	if user.UserType != UserType.Admin {
		res.Code = ErrNo.PermDenied
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	res.Code = VerifyCodeDao.DeleteVerifyCode(tx, request.VerifyCode)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
