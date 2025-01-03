package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main()  {
	fmt.Println("dir:",cwd)
}


func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working Directory = %s", cwd)
}
