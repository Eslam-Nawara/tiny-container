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

	switch os.Args[1] {
	case "run":
		initContainer()
		container.Run(os.Args[2:])
	case "child":
		container.Child(os.Args[2], os.Args[3:])
	default:
		panic("invalid command")
	}
}

func initContainer() {
	if _, err := os.Stat("/tmp/rootfs"); os.IsNotExist(err) {
		script, err := exec.Command("curl", "https://raw.githubusercontent.com/Eslam-Nawara/tiny-container/main/install.sh").Output()
		if err != nil {
			panic(err)
		}
		exec.Command("bash", "-c", string(script)).Run()
	}
}
