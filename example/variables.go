package example

import "fmt"

func Variables() {
	var a = "initial"
	var b, c int = 1, 2

	var d = true

	var e int

	f := "short"

	fmt.Println("a: ", a, "b: ", b, "c: ", c, "d: ", d, "e: ", e, "f: ", f)
}