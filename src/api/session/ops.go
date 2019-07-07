package session
import(
	"time"
	"sync"
	"api/defs"
	"api/dbops"
	"api/utils"
)

//支持并发读写Map,大量数据存储
var sessionMap *sync.Map
func init(){
	sessionMap =&sync.Map{}
}
func nowInMilli()int64{
	return time.Now().UnixNano()/1000000
}
func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
func LoadSessionsFromDB(){
	r,err:=dbops.RetrieveAllSessions()//r  ={int ,session}
	if err!=nil{
		return
	}
	r.Range(func(k,v interface{}) bool{
		ss:=v.(*defs.SimpleSession)//强转 v 至 *defs.SimpleSession
		sessionMap.Store(k,ss)//加载到全局 sessionMap
		return true
	})
}
func ClearAllSessions()error{
	if err:=dbops.ClearSessions();err!=nil{
		return err
	}
	return nil
}
func GenerateNewSessionId(un string)string{
	id,_:=utils.NewUUID()
	ct:=nowInMilli()
	ttl:=ct+30*60*1000 //30min
	ss:=&defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss) //缓存
	dbops.InsertSession(id,ttl,un) //数据库
	return id
}
//是否过期/存在
func IsSessionExpired(sid string)(string,bool){
	ss,ok:=sessionMap.Load(sid)
	if ok{
		ct:=nowInMilli()
		if ss.(*defs.SimpleSession).TTL <ct{
			deleteExpiredSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "" ,true
}