package main

import (
	"os"
	"os/exec"

	"github.com/Eslam-Nawara/tinycontainer"
)

func main() {
	if len(os.Args) < 2 {
		panic("Too few arguments")
	}

	cmd := exec.Command("./install.sh")
	cmd.Run()

	switch os.Args[1] {
	case "run":
		tinycontainer.Run(os.Args[2:])
	case "child":
		tinycontainer.Child(os.Args[2], os.Args[3:])
	default:
		panic("invalid command")
	}
}
