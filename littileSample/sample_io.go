package main 

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
	"strings"
)

// echo "hello go" | base64
const data = `
aGVsbG8gZ28K`

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecode(base64.StdEncoding, r)
	r, _ = gzip.NewReader(r)
	io.Copy(os.Stdout, r)
}