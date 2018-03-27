package main

import "testing"

func TestItRelaysFileChangesToCommandRunner(t *testing.T) {
	fileChangeMonitor := make(chan bool)
	commandRunnerChannel := make(chan string)

	go func() {
		relayFileChangesToCommandRunner(fileChangeMonitor, commandRunnerChannel)
	}()

	fileChangeMonitor <- true
	received := <-commandRunnerChannel

	if received != "restart" {
		t.Fail()
	}
}
