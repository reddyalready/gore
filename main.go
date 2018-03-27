package main

import (
	"log"
	"os"
)

func main() {

	filesChanged := watchFilesystem()
	commandRunner, err := startRunner(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	relayFileChangesToCommandRunner(filesChanged, commandRunner)
}

func relayFileChangesToCommandRunner(filesChanged <-chan bool, commandRunner chan<- string) {
	for {
		if <-filesChanged {
			commandRunner <- "restart"
		}
	}
}
