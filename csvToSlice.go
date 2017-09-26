package main

import (
	"bufio"
	"os"
	"strings"
)

// converts csv file to a slice taking the column number as an input
// if the file has a header is true then the slice removes the header
func csvToSlice(column int, header bool, fileLocation string) []string {
	f, _ := os.Open(fileLocation)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var s []string

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line on commas.
		parts := strings.Split(line, ",")
		s = append(s, parts[column])
	}
	if header == true {
		s = append(s[:0], s[0+1:]...)
	}
	return s
}
