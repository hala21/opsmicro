package user

import (
	"github.com/go-log/log"
	"opsmicro/ywhub/basic/db"
	user "opsmicro/ywhub/user-srv/proto/user"
)

func (s *service) QueryUserByName(userName string) (resp *user.User, err error) {
	queryString := `select user_id,user_name,pwd from user where user_name = ? `

	conn := db.GetDB()
	resp = &user.User{}

	err = conn.QueryRow(queryString, userName).Scan(&resp.Id, &resp.Username, &resp.Password)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据库失败,Error : %s", err)
		return
	}
	return
}
