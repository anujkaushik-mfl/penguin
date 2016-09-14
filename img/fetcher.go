package img

import (
	"github.com/anujkaushik-mfl/penguin/config"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func FetchImageFromHTTPSource(sourceUrl string) *ImageInfo {

	log.Println("Fetching image from URL : ", sourceUrl)

	imageFilePath := config.IMAGE_STORAGE_TEMP + createImageFileName()
	file, err := os.Create(imageFilePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	httpResponse, err := http.Get(sourceUrl)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer httpResponse.Body.Close()

	sizeBytes, err := io.Copy(file, httpResponse.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("File downloaded susseccfully : size = ", sizeBytes, "B")

	imageInfo := NewImageInfo(imageFilePath)

	if time.Now().Second()%2 == 0 {
		log.Println("Processing Image : Yes")
		imageInfo = ChangeImage(imageInfo)
	} else {
		log.Println("Processing Image : No")
	}

	return imageInfo
}

func createImageFileName() string {
	fileName := "sat_image_" + time.Now().Format("2006_Jan_02_15_04_05") + "." + config.IMAGE_EXTENTION
	return fileName
}

func ChangeImage(imageInfo *ImageInfo) *ImageInfo {

	originalImageFile, _ := os.Open(imageInfo.absolutePath)
	defer originalImageFile.Close()
	originalImage, _, _ := image.Decode(originalImageFile)

	newImage := IncreaseImageBrightness(&originalImage)

	newImageFilePath := config.IMAGE_STORAGE_TEMP + createImageFileName()
	newImageFile, _ := os.Create(newImageFilePath)
	defer newImageFile.Close()

	jpeg.Encode(newImageFile, newImage, nil)

	return NewImageInfo(newImageFilePath)
}
