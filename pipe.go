package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")

	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Println("couldn't obtain the stdout pipe for command ", err)
	}

	if err := cmd0.Start(); err != nil {
		fmt.Println("The command can not be startup", err)
		return
	}
}
