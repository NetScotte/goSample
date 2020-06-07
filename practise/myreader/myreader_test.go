package myreader

import (
	"testing"
)

func TestReader(t *testing.T) {
	Sample_ioreader()
}

func TestCopy(t *testing.T) {
	Sample_iocopy()
}

func TestReadAll(t *testing.T) {
	Sample_ioreadall()
}

func TestReadFull(t *testing.T) {
	Sample_ioreadfull()
}

func TestReadmore(t *testing.T) {
	Sample_iomore()
}