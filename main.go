package main

import (
	"log"

	"strings"

	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {

	heffe := runner([]string{"ls", "-la"})

	filesChanged := watch()

	log.Println("Telling runner to start")
	heffe <- "restart"

	for {
		log.Println("Main waiting")
		select {
		case <-filesChanged:
			log.Println("File change detected, telling runner to restart")
			heffe <- "restart"
		case output := <-heffe:
			log.Println("Program ended - printing output")
			log.Println(output)
			break
		}
	}
}

func runner(cmdArgs []string) chan string {

	runnerio := make(chan string)

	go func() {
		for {
			reloadMessage := <-runnerio

			if reloadMessage == "restart" {
				fullPath, err := exec.LookPath(cmdArgs[0])

				if err != nil {
					log.Fatal(err)
				}

				cmd := exec.Command(fullPath, cmdArgs[1:]...)
				output, err := cmd.CombinedOutput()

				if err != nil {
					log.Fatal(err)
				}

				runnerio <- string(output)
			}
		}
	}()

	return runnerio
}

func watch() <-chan bool {
	log.Println("Watch routine starting new file watcher")

	watchChannel := make(chan bool)

	//Start watching files
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	//Add the current directory
	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	//Process events
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				changedFile := event.Name
				if strings.HasSuffix(changedFile, ".go") {
					log.Printf("File change detected: %s", changedFile)
					close(watchChannel)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return watchChannel
}
