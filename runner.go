package main

import (
	"os"
	"errors"
	"log"
	"os/exec"
)

func startRunner(commandToRun []string) (chan<- string, error) {

	if len(commandToRun) < 1 {
		return nil, errors.New("runner: no arguments passed")
	}

	reloadChannel := make(chan string)

	go func() {
		log.Printf("Starting process %s", commandToRun)
		process, err := runCommand(commandToRun)
		if err != nil {
			log.Printf("Run failed: %s", err)
		}

		for {
			reloadMessage := <-reloadChannel

			if reloadMessage == "restart" {
				if err := process.Kill(); err != nil {
					log.Printf("Kill failed: %s", err)
				}
				log.Printf("Restarting process %s", commandToRun)
				process, err = runCommand(commandToRun)
				if err != nil {
					log.Printf("Run failed: %s", err)
				}
			}
		}
	}()

	return reloadChannel, nil
}

func runCommand(commandToRun []string) (*os.Process, error) {
	programPath, err := exec.LookPath(commandToRun[0])
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(programPath, commandToRun[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	cmd.Run()

	return cmd.Process, err
}
