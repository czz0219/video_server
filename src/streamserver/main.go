package main
import(
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
type middleWareHandler struct{
	r *httprouter.Router
	l *ConnLimiter
}
func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler {
	m:=middleWareHandler{}
	m.r=r
	m.l =NewConnLimiter(cc)
	return m
}
func RegisterHandlers() * httprouter.Router{
	router:=httprouter.New()
	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}
//AOP
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if !m.l.GetConn(){
		sendErrorResponse(w,http.StatusTooManyRequests,"Too many request")
		return
	}
	m.r.ServeHTTP(w,r)//router透传
	defer m.l.ReleaseConn()
}
func main(){
	r:=RegisterHandlers()
	mh:=NewMiddleWareHandler(r,2)
	log.Println("lensting 9000")
	http.ListenAndServe(":9000",mh)
}