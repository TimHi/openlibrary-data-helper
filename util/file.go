package util

import (
	"io/ioutil"
	"os"
	"strings"
)

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadLines(filename string) (map[int]string, error) {
	lines := make(map[int]string)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	for n, line := range strings.Split(string(data), "\n") {
		lines[n] = line
	}

	return lines, nil
}
