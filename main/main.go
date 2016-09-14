package main

import (
	"github.com/anujkaushik-mfl/penguin/config"
	"github.com/anujkaushik-mfl/penguin/img"
	"github.com/anujkaushik-mfl/penguin/rest"
	"github.com/anujkaushik-mfl/penguin/scheduler"
	"log"
)

func main() {

	log.Println("==== INIT ====")

	channel := make(chan int)
	imageUpdateJob := new(ImageUpdateJob)
	imageUpdateJob.Init(img.NewImageTracker())
	var job scheduler.Job = imageUpdateJob
	scheduler := scheduler.NewScheduler(config.FETCH_PERIOD_SECONDS, channel, &job)
	go scheduler.Start()

	log.Println("Webserver : Started")
	rest.StartRestServer()
	log.Println("Webserver : Terminated")

	channel <- -1

	log.Println("==== END ====")
}

type ImageUpdateJob struct {
	imageTracker *img.ImageTracker
}

func (o *ImageUpdateJob) Init(imageTracker *img.ImageTracker) {
	o.imageTracker = imageTracker
}

func (o *ImageUpdateJob) Do() {
	log.Println("-----> Job Start.")
	imageInfo := img.FetchImageFromHTTPSource(config.URL_IMAGE_SOURCE)
	o.imageTracker.SubmitImage(imageInfo)
	log.Println("-----> Job End.")
}
