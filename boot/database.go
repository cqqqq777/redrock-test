package boot

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	g "redrock-test/global"
)

func DatabaseInit() {
	MysqlInit()
	RedisInit()
}

func MysqlInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		g.Config.Database.Mysql.Username,
		g.Config.Database.Mysql.Password,
		g.Config.Database.Mysql.Addr,
		g.Config.Database.Mysql.Port,
		g.Config.Database.Mysql.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		g.Logger.Fatal(err.Error())
	}
	g.Mdb = db
	g.Logger.Info("connect mysql successfully")
}

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", g.Config.Database.Redis.Addr, g.Config.Database.Redis.Port),
		Password: g.Config.Database.Redis.Password,
		DB:       g.Config.Database.Redis.DB,
		PoolSize: g.Config.Database.Redis.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		g.Logger.Fatal(err.Error())
	}
	g.Rdb = client
	g.Logger.Info("connect redis successfully")
}
