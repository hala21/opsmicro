package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"net/http"
	auth "opsmicro/ywhub/auth/proto/auth"
	us "opsmicro/ywhub/user-srv/proto/user"
	"time"
)

var (
	serviceClient us.UserService
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("ywhub.micro.v1.basic.srv.user", client.DefaultClient)

	authClient = auth.NewService("ywhub.micro.v1.srv.auth", client.DefaultClient)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	_ = r.ParseForm()

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Password == r.Form.Get("pwd") {
		response["success"] = rsp.Success

		// 干掉密码返回
		rsp.User.Password = ""
		response["data"] = rsp.User
		log.Logf("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Username,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Register
func Register(w http.ResponseWriter, r *http.Request) {

	//判断POST方法
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	_ = r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	email := r.Form.Get("email")
	nike := r.Form.Get("nike")

	// 1.确认用户名是否唯一)
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{UserName: r.Form.Get(username)})

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Printf("%v", rsp
	if rsp.User != nil {

	}

	// 假如不存在的话，信息写入数据库

	// 获取token信息跳转到dashbord

}
