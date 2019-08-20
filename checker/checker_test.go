package checker

import (
	"fmt"
	"log"
)

func ExampleAdd() {
	fmt.Println(Add(1, 2))

	// Output: 3
}

func ExampleCheck() {
	fixed, err := Check("아버지가방에들어가신다")

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(fixed)

	// Output: 아버지가 방에 들어가신다
}

func ExampleCheckFile() {
	fixed, err := CheckFile("testdata/test.txt")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(fixed)

	// Output: 왜 안돼
}
