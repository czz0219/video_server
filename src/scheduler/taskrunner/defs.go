package taskrunner

import (
	"os"
)

const (
	READY_TO_DISPATCH = "d" //任务分派，配送
	READY_TO_EXECUTE  = "e" //任务取出执行
	CLOSE             = "c"
	VIDEO_DIR         = `\videos\`
)

func GetCompletePath() string {
	gobin := os.Getenv("GOBIN")
	complete_pt := gobin + VIDEO_DIR
	return complete_pt
}

type controlChan chan string
type dataChan chan interface{}
type fn func(dc dataChan) error
