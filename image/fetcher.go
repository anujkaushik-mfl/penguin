package image

import (
	"io"
	"os"
	"net/http"
	"time"
	"log"
)

func FetchImageFromHTTPSource(sourceUrl string) ImageF{
	fileName := "sat_image_" + time.Now().Format(time.UnixDate) + ".jpeg";
	absoluteFilePath := IMAGE_STORAGE_TEMP + fileName;

	file, err := os.Create(absoluteFilePath);
	if(err != nil){
		log.Fatal(err);
		return;
	}
	defer file.Close()

	httpResponse, err := http.Get(sourceUrl);
	if(err != nil){
		log.Fatal(err);
		return;
	}
	defer httpResponse.Body.Close();

	sizeBytes, err := io.Copy(file, httpResponse.Body);
	if(err != nil){
		log.Fatal(err);
		return;
	}
	log.Println("file downloaded susseccfully : size = " + sizeBytes + "B");
}




