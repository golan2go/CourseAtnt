package main

import (
	"fmt"
)

func main() {
	fmt.Println(worker())
	possibleBug()
}

func possibleBug() {
	for i := 0; i < 3; i++ {
		defer func(j int) {
			fmt.Println(j)
		}(i)
	}
	fmt.Println("no bug")
}
func worker() error {
	db, err := connect("A")
	if err != nil {
		return err
	}
	defer disconnect(db)

	db2, err := connect("B")
	if err != nil {
		return err
	}
	defer disconnect(db2)
	fmt.Println("working")
	//	panic("oopsie :(")
	return nil
}

func connect(dsn string) (DB, error) {
	return DB(dsn), nil
}

func disconnect(db DB) {
	fmt.Printf("disconnect from %s\n", db)
}

type DB string // TODO
