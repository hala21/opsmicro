package user

import (
	"fmt"
	user "opsmicro/ywhub/user-srv/proto/user"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

// service 服务
type service struct {
}

type Service interface {
	QueryUserByName(userName string) (resp *user.User, err error)
	CreateNewUser(userName string, password string, email string) (resp *user.User, err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetServices] GetServices 未初始化")
	}
	return s, nil
}

// 初始化服务
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}
