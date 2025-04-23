package main

import "fmt"

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	ptrToNumber, ok := ptrToNumber.(*int)

	if ptrToNumber.(*int) != nil && ok {
		fmt.Println(*ptrToNumber.(*int))
	} else {
		fmt.Println("nil")
	}
}

func main() {
	v := 10
	printNumber(&v) // 10

	var pv *int
	printNumber(pv) // nil

	pv = &v
	printNumber(pv) // 10
}
