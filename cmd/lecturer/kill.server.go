package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatal(err)
	}
}

func killServer(pidFile string) error {
	// TODO: read file, convert to int, kill, remove file
	// os.Open, io/ioutil.ReadAll, strconv.Atoi
	file, err := os.Open(pidFile)
	if err != nil {
		return err
		// return fmt.Errorf("can't open %q: %w", pidFile, err)
	}

	defer file.Close()
	/*
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("can't delete %s - %w", pidFile, err)
			}
		}()
	*/

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	s := strings.TrimSpace(string(data)) // []byte -> string
	pid, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("bad pid in %q - %w", pidFile, err)
	}

	log.Printf("killing %d\n", pid)
	err = os.Remove(pidFile)
	if err != nil {
		return err
	}
	return nil
}
