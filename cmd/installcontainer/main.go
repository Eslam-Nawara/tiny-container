package main

import (
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("./install.sh")
	cmd.Run()

	time.Sleep(5 * time.Second)
}
