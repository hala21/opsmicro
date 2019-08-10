package db

import (
	"database/sql"
	"log"
	"opsmicro/ywhub/basic/config"
)

func initMysql() {
	var err error

	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxOpenConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

}
