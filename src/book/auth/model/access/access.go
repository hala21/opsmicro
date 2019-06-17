package access

import (
	"book/basic/redis"
	"fmt"
	r "github.com/go-redis/redis"
	"sync"
)

var (
	s  *service
	m  sync.RWMutex
	ca *r.Client
)

type service struct {
}

type Service interface {
	MakeAccessToken(subject *Subject) (ret string, err error)
	GetCachedAccessToken(subject *Subject) (ret string, err error)
	DelUserAccessToken(token string) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetServices未初始化！")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	ca = redis.GetRedis()
	s = &service{}
}
