package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//processors := make([]*FileProcessor, 0, 100)

	files, err := filepath.Glob("/Users/ig2258/Desktop/GoCourseResources/graphs/*")
	if err != nil {
		fmt.Println("Can't access files: ", err)
		return
	}

	messagesPerUser := make(map[int]int)

	for _, fileName := range files {
		processor := NewFileProcessor(fileName)
		//processors = append(processors, processor)
		processor.process()
		for user, count := range processor.messagesPerUser {
			messagesPerUser[user] += count
		}
	}

	winner := keyWithMaxValue(messagesPerUser)
	fmt.Println("The winner is", winner, "with", messagesPerUser[winner], "messages")

}

func keyWithMaxValue(m map[int]int) int {
	max := -1
	winner := -1
	for k, v := range m {
		if v > max {
			winner = k
			max = v
		}
	}
	return winner
}
