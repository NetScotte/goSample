package myregexp

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

func Match() {
	pattern := "liu"
	s := "liufy47"
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		log.Info("not match")
	} else {
		log.Infof("matched: %v", matched)
	}
}

func FindAll() {
	pattern := "[0-9]+"
	s := "liufy47aa26dd12"
	cregexp, err := regexp.Compile(pattern)
	if err != nil {
		log.Error(err)
	}
	digitSlice := cregexp.FindAllString(s, -1)
	log.Info(digitSlice)

}
