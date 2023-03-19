package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func ReadData(fileName string) string {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(fileContent)
}

func CountCharactersMap(str *string) *map[int]int {
	charMap := make(map[int]int, 123)

	for _, char := range *str {
		charMap[int(char)]++
	}

	return &charMap
}

func CountCharactersSlice(str *string) *[]int {
	chars := make([]int, 123)

	for _, char := range *str {
		chars[int(char)]++
	}

	return &chars
}

var ResMapBench *map[int]int
var ResSliceBench *[]int

func BenchmarkMap(b *testing.B) {
	// Benchmark the CountCharactersMap function
	fileContent := ReadData("./1984.txt")
	for i := 0; i < b.N; i++ {
		ResMapBench = CountCharactersMap(&fileContent)
	}
}

func BenchmarkSlice(b *testing.B) {
	fileContent := ReadData("./1984.txt")
	for i := 0; i < b.N; i++ {
		ResSliceBench = CountCharactersSlice(&fileContent)
	}
}
