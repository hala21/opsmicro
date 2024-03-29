package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	consulConfig            defaultConsulConfig
	mysqlConfig             defaultMysqlConfig
	redisConfig             defaultRedisConfig
	jwtConfig               defaultJwtConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] 已经初始化过了")
	}

	// 加载yaml，基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	pt := filepath.Join(appPath, "conf")
	_ = os.Chdir(appPath)

	// 找到application.yaml 文件
	if err := config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	// 找到需要引入的新配置
	if err := config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}
	log.Logf("[Init] 加载配置文件：path: %s %+v\n", pt+sp+"application.yml", profiles)

	// 开始导入新文件
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			log.Logf("[Init] 加载配置文件：ppath: %s\n", filePath)
			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		if err := config.Load(sources...); err != nil {
			panic(err)
		}

	}

	// 赋值
	config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	config.Get(defaultRootPath, "redis").Scan(&redisConfig)
	config.Get(defaultRootPath, "jwt").Scan(&jwtConfig)

	// 标记已经初始化
	inited = true

}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}

func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}

func GetJwtConfig() (ret JwtConfig) {
	return jwtConfig
}
