package main

import "fmt"

func main() {

	var s1 []string
	var s2 []string
	var s3 []string

	// Convert csv file to slices
	s1 = csvToSlice(0, true, "test.csv")
	s2 = csvToSlice(1, true, "test.csv")
	s3 = csvToSlice(2, true, "test.csv")

	// print slice arrays
	for i := range s1 {
		fmt.Println(s1[i])
		if i+1 == len(s1) {
			fmt.Println()
		}
	}
	for i := range s2 {
		fmt.Println(s2[i])
		if i+1 == len(s2) {
			fmt.Println()
		}
	}
	for i := range s3 {
		fmt.Println(s3[i])
		if i+1 == len(s3) {
			fmt.Println()
		}
	}
}
