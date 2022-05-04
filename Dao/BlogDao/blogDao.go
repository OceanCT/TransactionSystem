package BlogDao

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
			"create table if not exists Blog(" +
				"id int primary key auto_increment," +
				"blogInfo varchar(255) unique not null," +
				"creatorID int not null references User(id)" +
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
func InsertBlog(tx *sql.Tx, blog *Types.Blog) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("insert into Blog(blogInfo, creatorID) values(?,?);\n")
	res, err := tx.Exec(sqlStr, blog.BlogInfo, blog.CreatorID)
	if err != nil {
		fmt.Println("Error happened when inserting blog in function BlogDao.InsertBlog()")
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "Error 1406: Data too long") {
			return ErrNo.ParamInvalid
		}
		return ErrNo.UnknownError
	}
	blog.ID, err = res.LastInsertId()
	if err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func FindBlog(tx *sql.Tx, res *[]Types.Blog) ErrNo.ErrNo {
	if DBErr != nil {
		return ErrNo.UnknownError
	}
	sqlStr := fmt.Sprintf("select blogInfo,creatorID from Blog;\n")
	rows, err := tx.Query(sqlStr)
	if err != nil {
		fmt.Println("Error happened when finding blogs in function BlogDao.FindBlog()")
		fmt.Println(err)
		return ErrNo.UnknownError
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error happened when finding blogs in function BlogDao.FindBlog()")
			fmt.Println(err)
		}
	}(rows)
	for rows.Next() {
		var tmp Types.Blog
		err := rows.Scan(&tmp.BlogInfo, &tmp.CreatorID)
		if err != nil {
			fmt.Println("Error happened when finding blogs in function BlogDao.FindBlog()")
			fmt.Println(err)
			return ErrNo.UnknownError
		} else {
			*res = append(*res, tmp)
		}
	}
	return ErrNo.OK
}
