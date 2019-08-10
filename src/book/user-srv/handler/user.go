package handler

/*
用户模型服务很简单，就是像数据库获取用户信息被返回给调用者
*/
import (
	"context"
	"github.com/micro/go-micro/util/log"
	us "opsmicro/src/book/user-srv/model/user"
	s "opsmicro/src/book/user-srv/proto/user"
)

type Service struct{}

var (
	userService us.Service
)

func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	rsp.Success = true

	return nil
}
