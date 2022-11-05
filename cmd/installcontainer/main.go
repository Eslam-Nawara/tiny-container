package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("./install.sh")
	cmd.Run()
}
