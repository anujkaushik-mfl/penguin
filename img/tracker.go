package img

import (
	"github.com/anujkaushik-mfl/penguin/config"
	"log"
	"os"
	"path/filepath"
	"time"
)

/*
Notes
http://satellite.imd.gov.in/img/3Dasiasec_vis.jpg
*/

type ImageTracker struct {
	activeImageQueue []*ImageInfo
}

func NewImageTracker() *ImageTracker {
	o := new(ImageTracker)
	o.init()
	return o
}

func (o *ImageTracker) SubmitImage(imageInfo *ImageInfo) {

	log.Println("Submitting image to image-tracker.")

	newImageInfo := imageInfo

	qLen := len(o.activeImageQueue)

	if qLen < config.CAPACITY {
		log.Println("Queue below capacity. Adding without checking.")
		o.activeImageQueue = append(o.activeImageQueue, newImageInfo)
		return
	}

	lastImageFile := o.activeImageQueue[qLen-1]
	if lastImageFile.hash != newImageInfo.hash {
		log.Println("Image changed. Poping oldest image. Adding new image")
		os.Remove(o.activeImageQueue[0].absolutePath)
		o.activeImageQueue = o.activeImageQueue[1:qLen]
		o.activeImageQueue = append(o.activeImageQueue, newImageInfo)
	} else {
		/* discard the file */
		log.Println("Image un-changed. Doing nothing.")
		os.Remove(newImageInfo.absolutePath)
	}
}

func (o *ImageTracker) cleanUpStorage() {
	log.Println("(X) Deleting all files in storage.")
	fileNames, _ := filepath.Glob(config.IMAGE_STORAGE_TEMP + "*." + config.IMAGE_EXTENTION)
	for _, fileName := range fileNames {
		log.Println("(X) Deleting :", fileName)
		os.Remove(fileName)
	}
}

/* Unused */
func (o *ImageTracker) moveImageToActiveStorage(imageFile ImageInfo) {

	oldPath := imageFile.absolutePath

	fileName := "sat_image_" + time.Now().Format(time.UnixDate) + ".jpeg"
	newPath := config.IMAGE_STORAGE_ACTIVE + fileName

	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("File moved to active storage : " + newPath)
}

func (o *ImageTracker) init() {
	o.activeImageQueue = make([]*ImageInfo, 0)
	o.cleanUpStorage()
}
