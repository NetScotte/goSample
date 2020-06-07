package myosexec

import (
	"log"
	"fmt"
	"os/exec"
)

func Sample_lookpath() {
	postion, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("(exec.LookPath) ls position: ", postion)
}

func Sample_cmd() {
	cmd := exec.Command("ls", "/tmp")
	result, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(cmd.Output) ls /tmp\n%s", result)
}