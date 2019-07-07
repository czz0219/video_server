package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent:%d", i)
		}
		return nil
	}
	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received:%v", d)
			default:
				break forloop //select 类似switch，没有匹配项，即，dataChan没有数据，就到default选项里,即跳出。
				//因为先设定了Dispatcher任务，那么 dataChan里边肯定有值，不会出现刚进来就 break forloop退出
			}
		}
		return errors.New("Executor")
	}
	runner := NewRunner(30, false, d, e)
	log.Println("init Data slice elements:", len(runner.Data))
	go runner.StartAll()
	time.Sleep(time.Nanosecond) //确保在1纳秒左右后 子协程随主协程一起退出.
}
