package main

import (
	"fmt"

	"github.com/matthewmpalen/mochi/core"
)

func main() {
	entity := core.NewEntity("A", 100, 50, 5, 5, 5)
	fmt.Println(entity)
}
