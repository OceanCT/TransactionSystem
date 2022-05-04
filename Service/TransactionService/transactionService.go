package TransactionService

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Dao/BillDao"
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Dao/TransactionDao"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"database/sql"
)

var db *sql.DB
var DBErr error

func init() {
	db, DBErr = DBAccessor.MysqlInit()
}

func ShowTransactions(userid int64) Types.ShowTransactionResponse {
	var res Types.ShowTransactionResponse
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
	res.Code = TransactionDao.FindTransactionByUserID(tx, userid, &res.Data.Transactions)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func SetTransactionStatus(userid int64, request Types.SetTransactionStatusRequest) Types.SetTransactionStatusResponse {
	var res Types.SetTransactionStatusResponse
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
	res.Code = TransactionDao.SetTransactionStatus(tx, userid, request.TransactionID, request.TransactionStatus)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
func CreateTransaction(userid int64, request Types.CreateTransactionRequest) Types.CreateTransactionResponse {
	var res Types.CreateTransactionResponse
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
	res.Code = TransactionDao.InsertTransaction(tx, userid, request.To, request.TransactionInfo)
	if res.Code == ErrNo.OK {
		res.Code = BillDao.InsertBill(tx, request.Bill)
	}
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
