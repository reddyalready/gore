package main

import (
	"testing"
)

func TestItErrorsWhenArgumentsEmpty(t *testing.T) {
	runnerChannel, err := startRunner([]string{})
	if err == nil || runnerChannel != nil {
		t.Fail()
	}
}

func TestItReloadsProgramWhenSentRestartSignal(t *testing.T) {
	runnerChannel, err := startRunner([]string{"echo", "test"})
	if err != nil {
		t.Fail()
	}
	runnerChannel <- "restart"
	//TODO: Not sure how to test this yet
}
