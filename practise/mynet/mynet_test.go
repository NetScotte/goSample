package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	Sample_server("localhost:9091")
}


func TestClient1(t *testing.T) {
	Sample_client("localhost:9091")
}

func TestClient2(t *testing.T) {
	Sample_client("localhost:9091")
}