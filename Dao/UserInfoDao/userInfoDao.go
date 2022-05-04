package UserInfoDao

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
			"create table if not exists UserInfo(" +
				"id int primary key auto_increment," +
				"userID int not null references User(id)," +
				"birthday varchar(25)," +
				"deathdate varchar(25)," +
				"userinfo varchar(655)" +
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

func FindUserInfoByUserID(tx *sql.Tx, userid int64, userinfo *Types.UserInfo) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := "select birthday,deathdate,userinfo from UserInfo where userid=?;"
	err := tx.QueryRow(sqlStr, userid).Scan(&userinfo.Birthday, &userinfo.DeathDate, &userinfo.UserInfo)
	if err != nil {
		fmt.Println("Error happened when finding userinfo in function UserInfoDao.FindUserInfoByUserID()")
		fmt.Println(err)
		return ErrNo.UserNotExisted
	}
	return ErrNo.OK
}

func UpdateUserInfoByUserID(tx *sql.Tx, userid int64, userinfo *Types.UserInfo) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := "update UserInfo set birthday=?,deathdate=?,userinfo=? where userid=?;"
	_, err := tx.Exec(sqlStr, userinfo.Birthday, userinfo.DeathDate, userinfo.UserInfo, userid)
	if err != nil {
		fmt.Println("Error happened when updating userinfo in function UserInfoDao.UpdateUserInfoByUserID()")
		fmt.Println(err)
		return ErrNo.UserNotExisted
	}
	return ErrNo.OK
}
