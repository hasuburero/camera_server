package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: [file name]")
		return
	}

	filename := os.Args[1]
	cmd := exec.Command("libcamera-jpeg", "-t", "1", "-o", filename)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("exec libcamera-jpeg error")
		return
	}

	fmt.Println("capture succeeded")
	return
}
