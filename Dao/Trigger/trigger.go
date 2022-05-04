package Trigger

import (
	"HomeworkForDB/Dao/DBAccessor"
	"database/sql"
	"fmt"
)

var db *sql.DB
var DBErr error

func TriggerInit() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		tx, err := db.Begin()
		if err != nil {
			DBErr = err
			fmt.Println(err)
			return
		}
		targetSql := fmt.Sprintf(
			"create trigger userinfo before insert on User for each row begin\nselect max(id) from User into @id; if @id is null then set @id = 1;else set @id = @id + 1;end if;\ninsert into UserInfo(userID,birthday,deathdate,userinfo) values (@id,'','','');\nend;")
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
