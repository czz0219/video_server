package taskrunner

import (
	"errors"
	"log"
	"os"
	"scheduler/dbops"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) //可能存在重复读数据，写channel。
	if err != nil {
		log.Printf("Video clear dispatcher error:%v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}
func deleteVideo(vid string) error {
	err := os.Remove(GetCompletePath() + vid)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Deleting video error:%v", err)
		return err
	}
	return nil
}
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop //当读空 dataChan，表示任务已经完成，不再需要阻塞读取了，退出死循环。
		}
	}
	//https://gitee.com/newdas/video_server/tree/master/scheduler
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return nil
}
