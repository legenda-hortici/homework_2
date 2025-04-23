package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
func main() {
	/*
		Даже если err является nil, он не равен nil, потому что имеет тип *MyError,
		соответсвенно err хранит интерфейсный тип и не является nil, но имеет nil значение
	*/
	var err *MyError        // type: *MyError, value: nil
	fmt.Println(err == nil) // true
	errorHandler(err)

	err = &MyError{}        // type: *MyError, value: &{MyError!}
	fmt.Println(err == nil) // false
	errorHandler(err)
}
