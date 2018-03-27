package main

import (
	"testing"
	"os/exec"
)

func TestWatchChannelReturnsTrueWhenFileChanged(t *testing.T) {
	removeFile("watchChannelTest.tmp")
	watcherChannel := watchFilesystem()

	touchFile("watchChannelTest.tmp")
	filesChanged := <-watcherChannel

	if filesChanged != true {
		t.Fail()
	}

	removeFile("watchChannelTest.tmp")
}

func removeFile(name string)  {
	cmd := exec.Command("rm", name)
	cmd.Run()
}

func touchFile(name string) {
	cmd := exec.Command("touch", name)
	cmd.Run()
}
