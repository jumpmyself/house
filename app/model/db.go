package model

import (
	"fmt"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn 所有的数据库操作放在这里
var Conn *gorm.DB
var Rdb *redis.Client

func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "house", "a7pp7i6XCWRaK42L", "192.168.67.149", "house")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	Conn = conn
}

func NewRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.67.149:6379",
		Password: "",
		DB:       0,
	})
	Rdb = rdb

	//初始化session
	store, _ = redisstore.NewRedisStore(context.TODO(), Rdb)
	return

}

func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
	_ = Rdb.Close()
}
