package image

import (
	"container/list"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
Notes
http://satellite.imd.gov.in/img/3Dasiasec_vis.jpg
*/

const CAPACITY = 100
const LOG_TAG = "image.tracker"

type ImageTracker struct {
	activeImageQueue list.List
}

func NewImageTracker() ImageTracker {
	o := new(ImageTracker)
	o.init()
	return o
}

func (o ImageTracker) AddImageToActiveList(filePath string) {

	newImageFile := NewImageFile(filePath)

	last := o.activeImageQueue.Back()

	lastImageFile := ImageFile(last)

	/* If new image is different that last image then,
	1. queue : push new and eject the oldest.
	2. storage : move image from temporary to active storage.
	*/
	if lastImageFile.hash != newImageFile.hash {
		log.Println(LOG_TAG + " : Image changed. Adding.")
		o.activeImageQueue.PushBack(&newImageFile)
		o.activeImageQueue.Remove(o.activeImageQueue.Front())
		o.moveImageToActiveStorage(newImageFile)
	} else {
		/* TODO : delete temp file here or make an overall deletion strategy */
		log.Println(LOG_TAG + " : Image un-changed. Doing nothing.")
	}
}

func (o ImageTracker) moveImageToActiveStorage(imageFile ImageFile) {

	oldPath := imageFile.absolutePath

	fileName := "sat_image_" + time.Now().Format(time.UnixDate) + ".jpeg"
	newPath := IMAGE_STORAGE_ACTIVE + fileName

	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(LOG_TAG + " : file moved to active storage : " + newPath)
}

func (o ImageTracker) init() {
	o.activeImageQueue = list.New()
}
