package scheduler

import "github.com/anujkaushik-mfl/penguin/img"

type Job interface {
	Init(imageTracker *img.ImageTracker)
	Do()
}
