package entry

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	g "redrock-test/business/global"
)

func InitDataBase() {
	dsn := "root:sjadfkshdk@tcp(127.0.0.1:3306)/storehouse"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("init database failed  err:%v", err)
		return
	}
	g.Db = db
}
