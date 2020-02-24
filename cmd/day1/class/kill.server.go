package main

import (
	"log"
	"os"
)

func main() {

	if err := killServer(); err != nil {
		log.Fatal(err)
	}

}

const fileName = "/tmp/kill.server.pid"

func killServer() error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer closeFile(file)

	var b = make([]byte, 100)

	_, err = file.Read(b)
	if err != nil {
		return err
	}

	var pid = 0
	for i := range b {
		if b[i] == 0 {
			break
		}
		digit := int(b[i]) - '0'
		pid = pid*10 + digit
	}

	print(pid)
	return nil
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		print(err)
	}
}
