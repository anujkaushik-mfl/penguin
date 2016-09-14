package scheduler

import (
	"log"
	"time"
)

const MSG_SCHEDULER_STOP int = -1

type Scheduler struct {
	runForever    bool
	periodSeconds int
	inChannel     chan int
	job           *Job
}

func NewScheduler(periodSeconds int, inChannel chan int, job *Job) *Scheduler {
	o := new(Scheduler)
	o.periodSeconds = periodSeconds
	o.runForever = true
	o.inChannel = inChannel
	o.job = job
	return o
}

func (o Scheduler) Start() {
	log.Println("Scheduler : Started")
	runForever := true
	for runForever {
		select {
		case msg := <-o.inChannel:
			{
				if msg == MSG_SCHEDULER_STOP {
					log.Println("Stop recieved in channel")
					runForever = false
				} else {
					log.Println("Unrecognised message : %d : Ignoring.", msg)
				}
			}
		default:
			(*(o.job)).Do()
			time.Sleep(time.Duration(o.periodSeconds) * time.Second)
		}
	}
	log.Println("Scheduler : Terminated")
}
