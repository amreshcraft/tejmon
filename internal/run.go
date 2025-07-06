package internal

import (
	"fmt"
	"os"
	"os/exec"
)
func Run() {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()
	fmt.Println("🚀 App running...")
}