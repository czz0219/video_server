package dbops
import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"strconv"
	"api/defs"
	"log"
)
func InsertSession(sid string,ttl int64,uname string)error{
	ttlstr:=strconv.FormatInt(ttl,10)
	stmtIns,err:=dbConn.Prepare("insert into sessions (session_id,ttl,login_name) values (?,?,?)")
	if err!=nil{
		return err
	}
	_,err=stmtIns.Exec(sid,ttlstr,uname)
	if err!=nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}
func RetriveSession(sid string)(*defs.SimpleSession,error){
	ss:=&defs.SimpleSession{}
	stmtOut,err:=dbConn.Prepare("select ttl,login_name from sessions where session_id =?")
	if err!=nil{
		return nil,err
	}
	var ttl,uname string
	stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err!=nil &&err!=sql.ErrNoRows{
		return nil,err
	}
	if res,err:=strconv.ParseInt(ttl,10,64);err==nil{
		ss.TTL =res
		ss.Username = uname
	}else{
		return nil,err
	}
	defer stmtOut.Close()
	return ss,nil
}
func RetrieveAllSessions()(*sync.Map,error){
	m:=&sync.Map{}
	stmtOut,err:=dbConn.Prepare("select * from sessions")
	if err!=nil{
		log.Printf("%s",err)
		return nil,err
	}
	rows,err:= stmtOut.Query()
	if err!=nil{
		log.Printf("%s",err)
		return nil,err
	}
	for rows.Next(){
		var id ,ttlstr ,login_name string
		if err:=rows.Scan(&id,&ttlstr,&login_name);err!=nil{
			log.Printf("retruve sessions error:%s",err)
			break
		}
		if ttl,err1:=strconv.ParseInt(ttlstr,10,64);err1==nil{
			ss:=&defs.SimpleSession{Username:login_name,TTL:ttl}
			m.Store(id,ss)
			log.Printf(" session id:%s,ttl:%d",id,ss.TTL)
		}
	}
	//defer stmtOut.Close()
	return m,nil
}
func DeleteSession(sid string)error{
	stmtOut,err:=dbConn.Prepare("delete from sessions where session_id=?")
	if err!=nil{
		return err
	}
	if _,err:=stmtOut.Query(sid);err!=nil{
		return err
	}
	defer stmtOut.Close()
	return nil
}
func ClearSessions()error{
	stmtOut,err:=dbConn.Prepare("delete from sessions")
	if err!=nil{
		log.Printf("prepare delete fault:%s\n",err)
		return err
	}
	if _,err:=stmtOut.Exec();err!=nil{
		log.Printf("execute delete fault:%s\n",err)
		return err
	}
	return nil
}