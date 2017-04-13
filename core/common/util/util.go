package util

import (
	"fmt"
	"hash/crc32"
	"os"
	"regexp"
	"strings"
)

//JsonpToJSON modify jsonp string to json string
func JsonpToJSON(json string) string {
	if len(json) == 0 {
		return ""
	}
	start := strings.Index(json, "(")
	end := strings.LastIndex(json, ")")
	if start > -1 && end > -1 && start < end {
		return string(json[start:end])
	}

	return ""
}

//GetWorkDirectoryPath return the workk directory path
func GetWorkDirectoryPath() string {
	wd := os.Getenv("GOPATH")
	if wd == "" {
		panic("GOPATH is not setted in env.")
	}

	return wd
}

//IsDirectoryExists check that is directory exists
func IsDirectoryExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return fi.IsDir()
}

//IsFileExists check that is file exists
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return !fi.IsDir()
}

//IsNumber check that string expression type is number
func IsNumber(str string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(str)
}

//MakeHash genegrate hash value from string value
func MakeHash(str string) string {
	const IEEE = 0xedb88320
	var IEEETable = crc32.MakeTable(IEEE)
	hash := fmt.Sprintf("%x", crc32.Checksum([]byte(str), IEEETable))
	return hash
}
