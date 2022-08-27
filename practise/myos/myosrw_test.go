package myos

import (
	"testing"
)

func Test_rw(t *testing.T) {
	var content []byte
	Sample_readfile("test.text")
	Sample_writefile(content, "test.txt")
	Sample_filesystem()
}
