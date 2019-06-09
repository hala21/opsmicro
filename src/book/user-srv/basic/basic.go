package basic

import (
	"book/user-srv/basic/config"
	"book/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
