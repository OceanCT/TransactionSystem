package BillDao

import (
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
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
			"create table if not exists Bill(" +
				"id int primary key auto_increment," +
				"userID int not null references User(ID)," +
				"billMoney float not null," +
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

func FindBill(tx *sql.Tx, userid int64, res *[]Types.Bill) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select id, billMoney,status from Bill where userID = ?;")
	rows, err := tx.Query(sqlStr, userid)
	if err != nil {
		fmt.Println("Error happened when finding bills in function BillDao.FindBill()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error happened when finding bills in function BillDao.FindBill()")
			fmt.Println(err)
		}
	}(rows)
	for rows.Next() {
		var tmp = Types.Bill{UserID: userid}
		err := rows.Scan(&tmp.ID, &tmp.BillMoney, &tmp.Status)
		if err != nil {
			fmt.Println("Error happened when finding bills in function BillDao.FindBill()")
			fmt.Println(err)
			return ErrNo.UnknownError
		} else {
			*res = append(*res, tmp)
		}
	}
	return ErrNo.OK
}

func PayBill(tx *sql.Tx, billID int64) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf(
		"with tmp as (select userID,billMoney from Bill where Bill.id = ?)\nselect ID,money,billMoney from User join tmp where User.id = tmp.UserID;")
	var userid int64
	var money float64
	var billMoney float64
	err := tx.QueryRow(sqlStr, billID).Scan(&userid, &money, &billMoney)
	if err != nil {
		fmt.Println("Error happened when updating bill in function BillDao.Paying()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	if money < billMoney {
		return ErrNo.MoneyNotEnough
	}
	sqlStr = fmt.Sprintf("update User set money = money - ? where id = ?;")
	_, err = tx.Exec(sqlStr, billMoney, userid)
	if err != nil {
		fmt.Println("Error happened when updating bill in function BillDao.Paying()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	sqlStr = fmt.Sprintf("update Bill set status = 1 where id = ?;")
	_, err = tx.Exec(sqlStr, billID)
	if err != nil {
		fmt.Println("Error happened when updating bill in function BillDao.Paying()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}

func InsertBill(tx *sql.Tx, bill Types.Bill) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("insert into Bill(userID, billMoney, status) values(?, ?, ?)")
	_, err := tx.Exec(sqlStr, bill.UserID, bill.BillMoney, bill.Status)
	if err != nil {
		fmt.Println("Error happened when inserting bill in function BillDao.InsertBill()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
