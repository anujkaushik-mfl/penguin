package img

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
)

func CalculateFileHash(filePath string) string {
	log.Println("Calculating file hash.")
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filePath)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	hashString := hex.EncodeToString(hasher.Sum(nil))
	log.Println("Hash = ", hashString)
	return hashString
}
