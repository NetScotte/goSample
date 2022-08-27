package myos

// os标准包，是一个比较重要的包，顾名思义，主要是在服务器上进行系统的基本操作，如文件操作，目录操作，执行命令，信号与中断，进程，系统状态等等。
//	在os包下，有exec，signal，user三个子包。

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Sample_lookpath() {
	// 寻找ls的位置，类似which ls
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

func Sample_cmd_run() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
