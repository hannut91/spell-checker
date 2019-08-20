package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hannut91/spell-checker/checker"
)

func main() {
	filename := os.Args[1]

	fmt.Println("origin")
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	fixed, err := checker.CheckFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("fixed")
	fmt.Println(fixed)
}
