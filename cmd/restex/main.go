package main

import (
	"fmt"
	"os"

	"github.com/joshibrom/restex/internal/model"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("Invalid number of arguments supplied")
	}
	resume, err := model.LoadFromFile(args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", resume)
}
