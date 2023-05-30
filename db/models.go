package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	driverName := "mysql"
	dataSourceName := "root" + ":" + password + "@" + "tcp" + "(" + dbUrl + ")" + "/" + "spider_show"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("数据库连接错误: %s", err.Error())
	}
	return db, nil
}
