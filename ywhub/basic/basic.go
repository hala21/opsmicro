package basic

import (
	"book/basic/config"
	"book/basic/db"
	"book/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
