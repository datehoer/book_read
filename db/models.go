package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	driverName := "mysql"
	dataSourceName := "root" + ":" + "qq125638" + "@" + "tcp" + "(" + "124.221.222.201:3306" + ")" + "/" + "spider_show"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("数据库连接错误: %s", err.Error())
	}
	return db, nil
}
