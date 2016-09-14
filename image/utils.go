package image

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"encoding/hex"
)

func CalculateFileHash(filePath string) string {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filePath)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	return (hex.EncodeToString(hasher.Sum(nil)));
}
