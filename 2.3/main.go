package main

import "fmt"

func main() {
	var foo []int // len = 0, cap = 0
	var bar []int // len = 0, cap = 0

	foo = append(foo, 1) // foo: len = 1, cap = 1 [1]
	foo = append(foo, 2) // foo: len = 2, cap = 2 [1 2]
	foo = append(foo, 3) // foo: len = 3, cap = 4 [1 2 3]
	bar = append(foo, 4) // bar: len = 4, cap = 4 [1 2 3 4] foo: len = 3, cap = 4 [1 2 3]
	foo = append(foo, 5) // foo: len = 5, cap = 8 [1 2 3 5] bar: len = 5, cap = 8 [1 2 3 5]

	fmt.Println(foo, bar) // [1 2 3 5] [1 2 3 5]
}
