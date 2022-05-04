package ThumbDao

import (
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Types"
	"HomeworkForDB/Types/Const/ErrNo"
	"HomeworkForDB/Types/Const/ThumbStatus"
	"database/sql"
	"fmt"
	"strings"
)

var db *sql.DB
var DBErr error

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
			"create table if not exists Thumb(" +
				"id int primary key auto_increment," +
				"blogID int not null references User(id)," +
				"userID int not null references User(id)," +
				"thumb int not null" +
				");")
		_, err = tx.Exec(targetSql)
		if err != nil {
			DBErr = err
			fmt.Println(err)
		} else {
			targetSql = fmt.Sprintf("create unique index thumb_index on Thumb(blogID,userID);")
			_, err = tx.Exec(targetSql)
		}
		err = tx.Commit()
		if err != nil {
			DBErr = err
			fmt.Println(err)
		}
	}
}

func Thumb(tx *sql.Tx, blogID int64, userid int64, status ThumbStatus.ThumbStatus) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf(
		"insert into Thumb (blogID, userID, thumb) value (?,?,?) on duplicate key update thumb = ?;")
	_, err := tx.Exec(sqlStr, blogID, userid, status, status)
	if err != nil {
		fmt.Println("Error happened when thumbing in function ThumbDao.thumbing()")
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

func FindBlogThumb(tx *sql.Tx, blogID int64, res *Types.FindBlogThumbResponse) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select sum(sum1),sum(sum2) from BlogThumb where blogID= ?;")
	err := tx.QueryRow(sqlStr, blogID).Scan(&res.Data.ThumbSum1, &res.Data.ThumbSum2)
	if err != nil {
		res.Data.ThumbSum1 = 0
		res.Data.ThumbSum2 = 0
		fmt.Println("Error happened when finding blog thumb in function ThumbDao.FindBlogThumb()")
		fmt.Println(err)
	}
	return ErrNo.OK
}
