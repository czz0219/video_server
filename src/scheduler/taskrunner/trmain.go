package taskrunner

import (
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}
func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C: //C代表一个channel，每到定时时间就可以从里边取出一个tickets
			go w.runner.StartAll()
		}
	}
}
func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(5, r)
	go w.startWorker()
	//other timer and runner
}
