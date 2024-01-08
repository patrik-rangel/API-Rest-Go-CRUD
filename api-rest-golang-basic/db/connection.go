package db

import (
	"database/sql"
	"fmt"

	"api-rest-golang-basic/configs" 
	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("mysql", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
