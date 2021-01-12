package main

import (
	"os"
	"strconv"
)

func fillString(returnString string, toLength int) string {
	for len(returnString) < toLength {
		returnString = returnString + ":"
	}
	return returnString
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func intToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}
