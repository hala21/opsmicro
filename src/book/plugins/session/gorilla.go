package session

import (
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"net/http"
	"strings"
	"time"
)

var (
	sessionIDNamePrefix = "session-id-"
	store               *sessions.CookieStore
)

func Init() {
	store = sessions.NewCookieStore([]byte("OnNUU5RUr6Ii2HMI0d6E54bXTS52tCCL"))
}

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var sId string
	// 判断cook里面是否包含session
	for _, cook := range r.Cookies() {
		if strings.Index(cook.Name, sessionIDNamePrefix) == 0 {
			sId = cook.Name
			break
		}
	}
	// 假如session为空，就设置session
	if sId == "" {
		sId = sessionIDNamePrefix + uuid.New().String()
	}
	// 忽略错误，因为Get方法会一直返回错误
	ses, _ := store.Get(r, sId)
	// 没有ID说明是新session
	if ses.ID == "" {
		cookie := &http.Cookie{Name: sId, Value: sId, Path: "/", Expires: time.Now().Add(30 * time.Second), MaxAge: 0}
		http.SetCookie(w, cookie)
		ses.ID = sId
		ses.Save(r, w)

	}

	return ses

}
