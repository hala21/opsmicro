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

type UserInfo struct {
	UserName        string `json:"userName,omitempty" `
	FullName        string `json:"fullName,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password, omitempty"`
	ConfirmPassword string `json:"confirmPassword, omitempty"`
	Terms           string `json:"terms, omitempty"`
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
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

type reg struct {
	fullName        string
	password        string
	email           string
	confirmPassword string
	terms           bool
}

// Register
func Register(w http.ResponseWriter, r *http.Request) {

	//判断POST方法
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	//r.ParseForm()
	//fmt.Printf("%v", r.PostForm
	//for k, v :=range(){
	//	print(k,  v)
	//}

	//var reg reg
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf)
	//println(string(buf[0:n]))
	var userInfo UserInfo
	err := json.Unmarshal(buf[0:n], &userInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(userInfo)

	userName := r.Form.Get("fullName")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	//nike := r.Form.Get("nike")
	print(email)

	// 1.确认用户名是否唯一)
	rspUser, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{UserName: userName})

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Printf("%v", rspUser.Success)
	if rspUser.User != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		// 返回结果
		response := map[string]interface{}{
			"ref":     time.Now().UnixNano(),
			"message": "用户已经存在，请重新考虑个靓名",
		}
		// 返回JSON结构
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 假如不存在的话，信息写入数据库
	respCreate, err := serviceClient.CreateNewUser(context.TODO(), &us.Request{UserName: userName, UserPwd: password, Email: email})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	response["success"] = true

	// 干掉密码返回
	respCreate.User.Password = ""
	response["data"] = respCreate.User

	// 生成token
	rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
		UserId:   respCreate.User.Id,
		UserName: respCreate.User.Username,
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

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 获取token信息跳转到dashbord

}
