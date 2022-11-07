package main

import (
	"os"
	"os/exec"

	"github.com/Eslam-Nawara/tiny-container/internal/container"
)

func main() {
	if len(os.Args) < 2 {
		panic("Too few arguments")
	}

	initContainer()
	switch os.Args[1] {
	case "run":
		container.Run(os.Args[2:])
	case "child":
		container.Child(os.Args[2], os.Args[3:])
	default:
		panic("invalid command")
	}
}

func initContainer() {
	if _, err := os.Stat("/tmp/rootfs"); os.IsNotExist(err) {
		exec.Command("wget", "https://github.com/Eslam-Nawara/tiny-container/raw/main/install.sh").Output()
		os.Chmod("install.sh", 0777)
		exec.Command("bash", "-c", "./install.sh").Output()
		os.Remove("install.sh")
	}
}
