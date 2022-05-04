package BillService

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Dao/BillDao"
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"database/sql"
)

var db *sql.DB
var DBErr error

func init() {
	db, DBErr = DBAccessor.MysqlInit()
}

func ShowBill(userid int64) Types.ShowBillResponse {
	var res Types.ShowBillResponse
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
	res.Code = BillDao.FindBill(tx, userid, &res.Data.Bills)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func PayBill(request Types.PayBillRequest) Types.PayBillResponse {
	var res Types.PayBillResponse
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
	res.Code = BillDao.PayBill(tx, request.BillID)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
