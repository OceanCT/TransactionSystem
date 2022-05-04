package DBView

import (
	"HomeworkForDB/Dao/DBAccessor"
	"database/sql"
	"fmt"
)

var db *sql.DB
var DBErr error

func ViewInit() {
	db, DBErr = DBAccessor.MysqlInit()
	if DBErr == nil {
		tx, err := db.Begin()
		if err != nil {
			DBErr = err
			fmt.Println(err)
			return
		}
		targetSql := fmt.Sprintf(
			"create view BlogThumb as\nselect IFNULL(blogID1,blogID2) as blogID, IFNULL(sum1,0) as sum1, IFNULL(sum2,0) as sum2\nfrom (with\ntmp1 as(select blogID as blogID1,sum(thumb) as sum1  from (select blogID,thumb from Thumb where thumb=1) as tmp1 group by tmp1.blogID),\ntmp2 as(select blogID as blogID2,sum(thumb) as sum2  from (select blogID,thumb from Thumb where thumb=-1) as tmp2 group by tmp2.blogID)\n(select * from tmp1 left join tmp2 on tmp1.blogID1=tmp2.blogID2) union (select * from tmp1 right join tmp2 on tmp1.blogID1=tmp2.blogID2)) as alias1;\n")
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
