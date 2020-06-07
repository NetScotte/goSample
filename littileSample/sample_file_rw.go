package main

import (
	"fmt"
	"os"
	"log"
)

func write2file(content, filename string){
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(content)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// todo: 如何有效得读取文件内容, 
func readfile(filename string) (content string){
	b := []byte{1,2,3,4,5}
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	num, err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("read length: ", num)
	content = string(b)
	return content
}

func main() {
	// write2file("hello file", "file.txt")
	content := readfile("file.txt")
	fmt.Println(content)
}