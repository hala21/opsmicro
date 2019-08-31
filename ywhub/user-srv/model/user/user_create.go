package user

import (
	"context"
	"opsmicro/ywhub/basic/db"
	user "opsmicro/ywhub/user-srv/proto/user"
)

func (s *service) CreateNewUser(userName string, password string, email string) (resp *user.User, err error) {
	insertString := `insert into user (user_id,user_name,pwd,email,create_date,update_date) value(?, ?, ?, ?, ?)`

	conn := db.GetDB()
	resp = &user.User{}

	stmt, err := conn.ExecContext(context.TODO(), insertString)
	err = conn.QueryRow(queryString, userName).Scan(&resp.Id, &resp.Username, &resp.Password)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据库失败,Error : %s", err)
		return
	}
	return
}
