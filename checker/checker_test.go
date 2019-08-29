package checker

import (
	"fmt"
	"log"
)

func ExampleCheck() {
	fixed, err := Check("아버지가방에들어가신다")

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(fixed)

	// Output: 아버지가 방에 들어가신다
}
