package mycrypto

import (
	"crypto/md5"
)

func SampleMD5() string {
	s := "my string: 中国"
	d := md5.Sum([]byte(s))
	return string(d[:])
}
