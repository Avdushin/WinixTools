package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
