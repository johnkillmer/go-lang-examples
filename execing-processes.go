package main

import "syscall"
import "os"
import "os/exec"
import "fmt"

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		fmt.Println(lookErr)
	}

	args := []string{"ls", "-l", "-a", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		fmt.Println(execErr)
	}

}
