package main

import (
	"log"
	"github.com/fsnotify/fsnotify"
)

func watchFilesystem() <-chan bool {
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

	//Process FS events
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				changedFile := event.Name
				log.Printf("File change detected: %s", changedFile)
				watchChannel <- true
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return watchChannel
}
