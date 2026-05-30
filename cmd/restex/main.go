package main

import (
	"fmt"
	"os"

	"github.com/joshibrom/restex/internal/model"
	"github.com/joshibrom/restex/internal/render"
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

	tex, err := render.Render(resume)
	if err != nil {
		panic(err)
	}

	fmt.Println(tex)
}
