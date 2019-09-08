package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"opsmicro/ywhub/user-srv/model/user"
	s "opsmicro/ywhub/user-srv/proto/user"
)

var (
	userService user.Service
)

type User struct{}

func Init() {
	var err error
	userService, err = user.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *User) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	return nil
}

func (e *User) CreateNewUser(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.CreateNewUser(req.UserName, req.UserPwd, req.Email)
	if err != nil {
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	return nil
}

//// Call is a single request handler called via client.Call or the generated client code
//func (e *User) Call(ctx context.Context, req *Request, rsp *Response) error {
//	log.Log("Received User.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}
//
//// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *User) Stream(ctx context.Context, req *StreamingRequest, stream User_StreamStream) error {
//	log.Logf("Received User.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Logf("Responding: %d", i)
//		if err := stream.Send(&StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *User) PingPong(ctx context.Context, stream User_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Logf("Got ping %v", req.Stroke)
//		if err := stream.Send(&Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}
