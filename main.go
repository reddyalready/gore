package main

import (
	"log"

	"strings"

	"os/exec"

	"os"

	"errors"

	"github.com/fsnotify/fsnotify"
)

func main() {

	filesChanged := fsWatch()
	heffe, err := runner(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	for {
		if <-filesChanged {
			heffe <- "restart"
		}
	}
}

func runner(cmdArgs []string) (chan<- string, error) {

	if len(cmdArgs) < 1 {
		return nil, errors.New("runner: no arguments passed")
	}

	runnerIO := make(chan string)

	go func() {
		log.Printf("Starting process %s", cmdArgs)
		proc := run(cmdArgs)

		for {
			reloadMessage := <-runnerIO

			if reloadMessage == "restart" {
				if err := proc.Kill(); err != nil {
					log.Printf("Kill failed: %s", err)
				}
				log.Printf("Restarting process %s", cmdArgs)
				proc = run(cmdArgs)
			}
		}
	}()

	return runnerIO, nil
}

func run(cmdArgs []string) *os.Process {
	programPath, err := exec.LookPath(cmdArgs[0])
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(programPath, cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return cmd.Process
}

func fsWatch() <-chan bool {
	log.Println("Starting filesystem watch on current directory")

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
					watchChannel <- true
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return watchChannel
}
