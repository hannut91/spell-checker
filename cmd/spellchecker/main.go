package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hannut91/spell-checker/checker"
	"github.com/hannut91/spell-checker/differ"
)

func main() {
	origin := os.Args[1]

	fixed, err := checker.Check(origin)
	if err != nil {
		fmt.Println(err)
		return
	}

	diff := differ.Diff(origin, strings.ReplaceAll(fixed, "<br>", "\n"))
	result := strings.ReplaceAll(diff, "\x1b[32m \x1b[0m", "\x1b[32;3m \x1b[0m")

	fmt.Print(result)
}
