package mycrypto

import (
	"crypto/md5"
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func SampleMD5() string {
	s := "my string: 中国"
	d := md5.Sum([]byte(s))
	return string(d[:])
}

func SampleCert() {
	certPath := "/Users/netscotte/cert/no_ca_cert/easy.crt"
	certKey := "/Users/netscotte/cert/no_ca_cert/easy.key"
	fp, err := os.Open(certPath)
	if err != nil {
		log.Fatalln(err)
	}
	certByte, err := io.ReadAll(fp)
	if err != nil {
		log.Fatalln(err)
	}

	fp, err = os.Open(certKey)
	if err != nil {
		log.Fatalln(err)
	}
	keyByte, err := io.ReadAll(fp)
	if err != nil {
		log.Fatalln(err)
	}

	cert, err := tls.X509KeyPair(certByte, keyByte)
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listen, err := tls.Listen("tcp", "localhost:8080", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	_ = listen
}
