package main

import (
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

//Read a file that is compressed with BZIp2
//Expand it
//Calculate the sha256
//Return ths sha value
func main() {
	fmt.Println(FileHash("~/Desktop/GoCourseResources/taxi/taxi-01.csv.bz2"))
}

func FileHash(path string) (string, error) {
	fileHandle, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer closeFile(fileHandle)

	zipReader := bzip2.NewReader(fileHandle)

	hash := sha256.New()

	_, err = io.Copy(hash, zipReader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func closeFile(fileHandle *os.File) {
	_ = fileHandle.Close()
}
