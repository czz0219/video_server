package main
import (
	"api/session"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//github 下载三方库: go get github.com/package_names
//例如: go get github.com/julienschmidt/httprouter
type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

//通过该方法实现AOP设计模式，切点切面
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}


func RegisterHandlers() *httprouter.Router {
	router:=httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}
func Prepare() {
	session.ClearAllSessions()
	session.LoadSessionsFromDB()//将数据库session加载到sessionMap

}
func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	log.Println("lensting 8000")
	http.ListenAndServe(":8000", mh)
}
