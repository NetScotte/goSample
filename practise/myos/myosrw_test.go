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

func TestFileOperator(t *testing.T) {
	filePath := "/Users/aa/a.txt"
	err := FileOperator(filePath)
	if err != nil {
		t.Error(err)
	}
}
