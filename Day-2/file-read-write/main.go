package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileRead1() error{
	bytes, err := ioutil.ReadFile("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-2/file-read-write/blah.txt")

	if err != nil {
		log.Fatal("Error: while reading file", err)

		fmt.Println("sdvsdfsdfsdfsdf")
		return err
	}

	fmt.Println(string(bytes))
}

func fileRead2() {
	file, err := os.Open("/Users/manishjindal/go/src/gitlab.com/google/go/training/classroom/Day-2/file-read-write/blah.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("Error: while reading file", err)
	}

	scanner := bufio.NewScanner(file)

	b := scanner.Scan()
	fmt.Println(b)
	fmt.Println(scanner.Text())

	b = scanner.Scan()
	fmt.Println(b)
	fmt.Println(scanner.Text())

	b = scanner.Scan()
	fmt.Println(b)
	fmt.Println(scanner.Text())

}

func main() {
	// fileRead1()
	fileRead2()
}
