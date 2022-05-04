package UserService

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Dao/UserDao"
	"HomeworkForDB/Dao/UserInfoDao"
	"HomeworkForDB/Dao/VerifyCodeDao"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"database/sql"
	"fmt"
)

var db *sql.DB
var DBErr error

func init() {
	db, DBErr = DBAccessor.MysqlInit()
}
func RegisterUser(request Types.RegisterUserRequest) Types.RegisterUserResponse {
	var res Types.RegisterUserResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	user := request.User
	fmt.Println(user)
	user.Money = 0
	verifyCode := request.VerifyCode
	if user.Username == "" {
		res.Code = ErrNo.UsernameNull
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	tx, err := db.Begin()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	res.Code = VerifyCodeDao.CheckVerifyCode(tx, user.UserType, verifyCode)
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	res.Code = UserDao.InsertUser(tx, &user)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func Login(request Types.LoginRequest) Types.LoginResponse {
	var res Types.LoginResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var user = Types.User{
		Username: request.Username,
	}
	tx, err := db.Begin()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	res.Code = UserDao.FindUserByUsername(tx, &user)
	fmt.Println(user)
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	if request.Password != user.Password {
		res.Code = ErrNo.WrongPassword
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var userinfo Types.UserInfo
	res.Code = UserInfoDao.FindUserInfoByUserID(tx, user.UserID, &userinfo)
	if res.Code == ErrNo.OK {
		res.Data.User = user
		res.Data.Userinfo = userinfo
	}
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Userinfo = Types.UserInfo{}
		res.Data.User = Types.User{}
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func WhoAmI(userid int64) Types.WhoAmIResponse {
	var res Types.WhoAmIResponse
	if DBErr != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var user = Types.User{
		UserID: userid,
	}
	fmt.Println(userid)
	tx, err := db.Begin()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	}
	res.Code = UserDao.FindUserByID(tx, &user)
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	var userinfo Types.UserInfo
	res.Code = UserInfoDao.FindUserInfoByUserID(tx, user.UserID, &userinfo)
	if res.Code == ErrNo.OK {
		res.Data.User = user
		res.Data.Userinfo = userinfo
	}
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
		res.Data.Userinfo = Types.UserInfo{}
		res.Data.User = Types.User{}
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func UpdateUser(userid int64, request Types.ChangeUserRequest) Types.ChangeUserResponse {
	var res Types.ChangeUserResponse
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
	res.Code = UserDao.UpdateUserByID(tx, userid, request.User)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func UpdateUserInfo(userid int64, request Types.ChangeUserInfoRequest) Types.ChangeUserInfoResponse {
	var res Types.ChangeUserInfoResponse
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
	res.Code = UserInfoDao.UpdateUserInfoByUserID(tx, userid, &request.UserInfo)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
