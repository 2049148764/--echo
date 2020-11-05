package utils

import (
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var Db *sqlx.DB // 全局的数据库操作

// init函数
func init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True"
	Db, err = sqlx.Open(`mysql`, dsn)
	if err != nil {
		panic("连接错误")
		os.Exit(1)
	}
	if err = Db.Ping(); err != nil {
		panic("运行错误")
	}
}