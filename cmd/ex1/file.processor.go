package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileProcessor struct {
	fileName        string
	messagesPerUser map[int]int
	successCount    int
	failCount       int
	err             error
}

func NewFileProcessor(fileName string) *FileProcessor {
	return &FileProcessor{
		fileName,
		make(map[int]int),
		0,
		0,
		nil,
	}
}

func (fp *FileProcessor) process() {
	fileHandle, err := os.Open(fp.fileName)
	if err != nil {
		fmt.Println("Can't access files: ", err)
		fp.err = err
	}
	defer fp.closeFileHandle(fileHandle)

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		receiver, err := extractReceiver(line)
		if err != nil {
			fp.failCount++
		} else {
			fp.messagesPerUser[receiver]++
			fp.successCount++
		}
	}
}

func (fp *FileProcessor) closeFileHandle(fileHandle *os.File) {
	err := fileHandle.Close()
	if err != nil {
		fp.err = err
	}
}

func extractReceiver(line string) (int, error) {

	fields := strings.Fields(line)
	if len(fields) != 3 {
		return -1, fmt.Errorf("can't parse")
	}

	_, err := strconv.Atoi(fields[0])
	if err != nil {
		return -1, fmt.Errorf("can't parse")
	}

	if fields[1] != "->" {
		return -1, fmt.Errorf("can't parse")
	}

	to, err := strconv.Atoi(fields[2])
	if err != nil {
		return -1, fmt.Errorf("can't parse")
	}

	return to, nil
}
