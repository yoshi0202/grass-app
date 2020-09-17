package services

import (
	"regexp"
)

func FindSvgTag(reg string, bytes []byte) []byte {
	return CreateRegexp(reg).Find([]byte(string(bytes)))
}

func FindAllRectTag(reg string, bytes []byte) [][]byte {
	return CreateRegexp(reg).FindAll([]byte(string(bytes)), -1)
}

func CreateRegexp(regString string) *regexp.Regexp {
	return regexp.MustCompile(regString)
}
