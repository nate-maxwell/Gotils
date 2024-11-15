package main

import (
	"fmt"
	"gotils/str"
)

func main() {
	foo := "theCamelAndTheSnake"
	bar := str.ConvertToSnakeCase(foo)
	baz := str.ConvertSnakeToPascale(bar)
	fmt.Println(baz)
}
