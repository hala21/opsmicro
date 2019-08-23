package basic

import (
	"opsmicro/ywhub/basic/config"
	"opsmicro/ywhub/basic/db"
	"opsmicro/ywhub/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
