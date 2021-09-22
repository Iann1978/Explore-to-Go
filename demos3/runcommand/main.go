//blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
package main

import (
	"os"
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
