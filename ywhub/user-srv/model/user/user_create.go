package user

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	"github.com/google/uuid"
	"opsmicro/ywhub/basic/db"
	user "opsmicro/ywhub/user-srv/proto/user"
)

func (s *service) CreateNewUser(userName string, password string, email string) (resp *user.User, err error) {
	insertString := `insert into user (user_id, user_name, pwd, email) value(?, ?, ?, ?)`
	queryString := `select user_id,user_name,pwd from user where user_name = ? `

	conn := db.GetDB()
	resp = &user.User{}

	userId, err := uuid.NewRandom()
	if err != nil {
		log.Logf("注册用户 生成UUID失败,error : %s", err)
	}

	stmt, _ := conn.Prepare(insertString)
	defer stmt.Close()

	ret, err := stmt.ExecContext(context.TODO(), insertString, userId, userName, email)
	if err != nil {
		log.Logf("[CreateNewUser] 注册用户失败,Error : %s", err)
		return
	}
	fmt.Printf("插入数据：%v", ret.LastInsertId())

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}

	err = conn.QueryRow(queryString, userName).Scan(&resp.Id, &resp.Username, &resp.Password)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据库失败,Error : %s", err)
		return
	}

	return resp, nil
}
