package db

import (
	"database/sql"
	"fmt"
	"github.com/go-log/log"
	"opsmicro/ywhub/basic/config"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	// 如果配置文件申明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	//

	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
