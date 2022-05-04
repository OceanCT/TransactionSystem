package BlogService

import (
	"HomeworkForDB/Config/ErrorInformation"
	"HomeworkForDB/Dao/BlogDao"
	"HomeworkForDB/Dao/DBAccessor"
	"HomeworkForDB/Dao/ThumbDao"
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

func AddBlog(userid int64, request Types.AddBlogRequest) Types.AddBlogResponse {
	var res Types.AddBlogResponse
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
	request.Blog.CreatorID = userid
	res.Code = BlogDao.InsertBlog(tx, &request.Blog)
	if res.Code != ErrNo.OK {
		res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
		return res
	}
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func ShowBlog() Types.ShowBlogResponse {
	var res Types.ShowBlogResponse
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
	res.Code = BlogDao.FindBlog(tx, &res.Data.Blogs)
	err = tx.Commit()
	if err != nil {
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func Thumb(userid int64, request Types.ThumbRequest) Types.ThumbResponse {
	var res Types.ThumbResponse
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
	res.Code = ThumbDao.Thumb(tx, request.BlogID, userid, request.Status)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}

func FindBlogThumb(request Types.FindBlogThumbRequest) Types.FindBlogThumbResponse {
	var res Types.FindBlogThumbResponse
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
	res.Code = ThumbDao.FindBlogThumb(tx, request.BlogID, &res)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		res.Code = ErrNo.UnknownError
	}
	res.Data.Message = ErrorInformation.ErrorInformation(res.Code)
	return res
}
