package main
import(
	"log"

)
type ConnLimiter struct{
	concurrentConn int
	bucket chan int
}
func NewConnLimiter(cc int)*ConnLimiter{
	return &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),
	}
}
func (cl *ConnLimiter)GetConn() bool{
//	log.Printf("connections num:%d",len(cl.bucket))
	if len(cl.bucket)>=cl.concurrentConn{
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}
func (cl *ConnLimiter)ReleaseConn(){
	_=<- cl.bucket
	//log.Printf("%d",c)
}