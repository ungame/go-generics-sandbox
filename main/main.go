package main

import (
	"fmt"
	"github.com/ungame/go-generics-sandbox/boolean"
)

func main() {
	v := true
	fmt.Println(boolean.IsTrue(v))

	v = false
	type Bool bool
	b := Bool(v)

	fmt.Println(boolean.IsTrue(b))
}
