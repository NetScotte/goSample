package main

import (
	"os"
	"bufio"
)

func SampleReader() {
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	SampleWriter("read str: " + str)
}

func SampleWriter(s string) {
	outputWriter := bufio.NewWriter(os.Stdout)
	outputWriter.WriteString(s)
	outputWriter.Flush()
}

func main() {
	SampleReader()
}