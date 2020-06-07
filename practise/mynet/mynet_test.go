package mynet

import (
	"testing"
)

func Test(t *testing.T) {
	go Sample_server("0.0.0.0:8181")
	Sample_client("localhost:8181")
	Sample_utils()
}