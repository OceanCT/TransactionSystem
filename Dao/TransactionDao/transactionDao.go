package TransactionDao

import (
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"HomeworkForDB/Types/Const/TransactionStatus"
	"database/sql"
	"fmt"
)

var db *sql.DB
var DBErr error

// connect to database and create table
func init() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		tx, err := db.Begin()
		if err != nil {
			DBErr = err
			fmt.Println(err)
			return
		}
		targetSql := fmt.Sprintf(
			"create table if not exists Transaction(" +
				"id int primary key auto_increment," +
				"fromID int not null references User(ID)," +
				"toID int not null references User(ID)," +
				"transactionInfo varchar(255) not null," +
				"status int not null" +
				");\n")
		_, err = tx.Exec(targetSql)
		if err != nil {
			DBErr = err
			fmt.Println(err)
		}
		err = tx.Commit()
		if err != nil {
			DBErr = err
			fmt.Println(err)
		}
	}
}
func FindTransactionByUserID(tx *sql.Tx, userid int64, res *[]Types.Transaction) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select id,transactionInfo,fromID,toID,status from Transaction where fromID = ? or toID =?;")
	rows, err := tx.Query(sqlStr, userid, userid)
	if err != nil {
		fmt.Println("Error happened when finding transactions in function TransactionDao.FindTransactionByUserID()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error happened when finding transactions in function TransactionDao.FindTransactionByUserID()")
			fmt.Println(err)
		}
	}(rows)
	for rows.Next() {
		var tmp = Types.Transaction{}
		err := rows.Scan(&tmp.ID, &tmp.TransactionInfo, &tmp.From, &tmp.To, &tmp.TransactionStatus)
		if err != nil {
			fmt.Println("Error happened when finding transactions in function TransactionDao.FindTransactionByUserID()")
			fmt.Println(err)
			return ErrNo.UnknownError
		} else {
			*res = append(*res, tmp)
		}
	}
	return ErrNo.OK
}
func SetTransactionStatus(tx *sql.Tx, userid int64, id int64, status TransactionStatus.TransactionStatus) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("update Transaction set status = ? where id = ? and toID = ?;")
	_, err := tx.Exec(sqlStr, status, id, userid)
	if err != nil {
		fmt.Println("Error happened when setting transaction status in function TransactionDao.SetTransactionStatus()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func InsertTransaction(tx *sql.Tx, from int64, to int64, transactionInfo string) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("insert into Transaction(fromID,toID,transactionInfo,status) values(?,?,?,1);")
	_, err := tx.Exec(sqlStr, from, to, transactionInfo)
	if err != nil {
		fmt.Println("Error happened when inserting transaction in function TransactionDao.InsertTransaction()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
