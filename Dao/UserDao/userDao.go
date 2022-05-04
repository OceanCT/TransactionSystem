package UserDao

import (
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
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
			"create table if not exists User(" +
				"id int primary key auto_increment," +
				"username varchar(25) unique not null," +
				"password varchar(25) not null," +
				"userType int not null," +
				"money double not null" +
				");\n")
		_, err = tx.Exec(targetSql)
		if err != nil {
			DBErr = err
			fmt.Println(err)
		} else {
			targetSql := fmt.Sprintf("create unique index username_index on User(username);")
			_, err = tx.Exec(targetSql)
		}
		err = tx.Commit()
		if err != nil {
			DBErr = err
			fmt.Println(err)
		}
	}
}

func InsertUser(tx *sql.Tx, user *Types.User) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("insert into User(username, password, userType,money) values(?,?,?,?);\n")
	res, err := tx.Exec(sqlStr, user.Username, user.Password, user.UserType, user.Money)
	if err != nil {
		fmt.Println("Error happened when inserting user in function UserDao.InsertUser()")
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "Error 1406: Data too long") {
			return ErrNo.ParamInvalid
		} else if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
			return ErrNo.UserHasExisted
		}
		return ErrNo.UnknownError
	}
	user.UserID, err = res.LastInsertId()
	if err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func FindUserByUsername(tx *sql.Tx, user *Types.User) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := "select id,password,userType,money from User where username=?;"
	err := tx.QueryRow(sqlStr, user.Username).Scan(&user.UserID, &user.Password, &user.UserType, &user.Money)
	if err != nil {
		fmt.Println("Error happened when finding user in function UserDao.FindUserByUserName()")
		fmt.Println(err)
		return ErrNo.UserNotExisted
	}
	return ErrNo.OK
}
func FindUserByID(tx *sql.Tx, user *Types.User) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := "select username,password,userType,money from User where id=?;"
	err := tx.QueryRow(sqlStr, user.UserID).Scan(&user.Username, &user.Password, &user.UserType, &user.Money)
	if err != nil {
		fmt.Println("Error happened when finding user in function UserDao.FindUserByID()")
		fmt.Println(err)
		return ErrNo.UserNotExisted
	}
	return ErrNo.OK
}
func UpdateUserByID(tx *sql.Tx, userid int64, user Types.User) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := "update User set password=?,money=? where id=?;"
	_, err := tx.Exec(sqlStr, user.Password, user.Money, userid)
	if err != nil {
		fmt.Println("Error happened when changing user by id in function UserDao.UpdateUserByID()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
