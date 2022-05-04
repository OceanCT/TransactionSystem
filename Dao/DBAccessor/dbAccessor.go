package DBAccessor

import (
	"HomeworkForDB/Config/DBConf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlInit() (*sql.DB, error) {
	dsn, err := DBConf.GetDsn()
	if err != nil {
		fmt.Println("Error happened when getting dsn in function MysqlInit()")
		fmt.Println(err)
		return nil, err
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error happened when connecting database in function MysqlInit()")
		fmt.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error happened when connecting database in function MysqlInit()")
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
