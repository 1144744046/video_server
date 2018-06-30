package task

import (
	"time"
)

type Worker struct {
	tick *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration,runner *Runner)*Worker{
    return &Worker{
    	time.NewTicker(interval*time.Second),
    	runner,
	}
}

func (work *Worker)StartWork(){
	for{
		select {
		case <-work.tick.C:
			go work.runner.StartAll()
		}
	}
}

func Start(){
	r:=NewRunner(3,true,VideoClearDispatch,VideoClearExecute)
	w:=NewWorker(3,r)
	go w.StartWork()
}