package VerifyCodeDao

import (
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"HomeworkForDB/Types/Const/UserType"
	"database/sql"
	"fmt"
	"strings"
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
			"create table if not exists VerifyCode(" +
				"id int primary key auto_increment," +
				"verifyCode varchar(25) not null," +
				"userType int not null);\n")
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

func CheckVerifyCode(tx *sql.Tx, userType UserType.UserType, verifyCode string) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select id from VerifyCode where userType=? and verifyCode=?")
	var res int
	err := tx.QueryRow(sqlStr, userType, verifyCode).Scan(&res)
	if err != nil {
		fmt.Println("Error happened when checking VerifyCode in function VerifyCodeDao.CheckingVerifyCode()")
		fmt.Println(err)
		return ErrNo.VerifyCodeNotValid
	}
	return ErrNo.OK
}

func InsertVerifyCode(tx *sql.Tx, verifyCode Types.VerifyCode) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("insert into VerifyCode(verifyCode,userType) values(?,?);\n")
	_, err := tx.Exec(sqlStr, verifyCode.VerifyCode, verifyCode.UserType)
	if err != nil {
		fmt.Println("Error happened when inserting verifyCode in function UserDao.InsertVerifyCode()")
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "Error 1406: Data too long") {
			return ErrNo.ParamInvalid
		}
		return ErrNo.UnknownError
	}
	if err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}

func FindAllVerifyCode(tx *sql.Tx, res *[]Types.VerifyCode) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select verifyCode,userType from VerifyCode;\n")
	rows, err := tx.Query(sqlStr)
	if err != nil {
		fmt.Println("Error happened when finding verifyCode in function FindAllVerifyCode()")
		fmt.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error happened when finding verifyCode in function FindAllVerifyCode()")
			fmt.Println(err)
		}
	}(rows)
	for rows.Next() {
		var tmp Types.VerifyCode
		err := rows.Scan(&tmp.VerifyCode, &tmp.UserType)
		if err != nil {
			fmt.Println("Error happened when finding verifyCode in function VerifyCodeDao.FindAllVerifyCode()")
			fmt.Println(err)
			return ErrNo.UnknownError
		} else {
			*res = append(*res, tmp)
		}
	}
	return ErrNo.OK
}

func DeleteVerifyCode(tx *sql.Tx, code string) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("delete from VerifyCode where verifyCode = ?\n")
	_, err := tx.Exec(sqlStr, code)
	if err != nil {
		fmt.Println("Error happened when deleting user by id in function VerifyCodeDao.DeleteVerifyCode()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
