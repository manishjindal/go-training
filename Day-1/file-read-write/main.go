package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileread1() {
	bytes, err := ioutil.ReadFile("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-1/file-read-write/blah.txt")

	if err != nil {
		log.Fatal("Failed to read the file", err)
	}

	fmt.Println(string(bytes))
}

func fileread2() {
	bytes, err := os.ReadFile("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-1/file-read-write/blah.txt")

	if err != nil {
		log.Fatal("Failed to read the file", err)
	}

	fmt.Println(string(bytes))
}

func fileread3() {
	file, err := os.Open("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-1/file-read-write/blah.txt")

	defer file.Close()

	if err != nil {
		log.Fatal("Failed to read the file", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	fileread1()
	fileread2()
	fileread3()
	/*
		file, err := os.Open("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-1/file-read-write/blah.txt")

		defer file.Close()

		if err != nil {
			log.Fatal("Failed to open file", err)
		}

		r := bufio.NewReader(file)

		file.Read()
		fmt.Println()
	*/
}
