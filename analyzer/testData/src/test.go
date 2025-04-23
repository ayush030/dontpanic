package testData

import (
	"fmt"
)

func ExportedTestFunctionWithPanic() {
	var a interface{} = 10

	switch v := a.(type) {
	case string:
		panic("string not expected ") // want "avoid use of panic"
	case int:
		fmt.Println("int is expected", v)
	default:
		panic("unknown type") // want "avoid use of panic"
	}
}

func unexportedTestFunctionWithPanic() {
	var a interface{} = 10

	switch v := a.(type) {
	case string:
		panic(fmt.Sprintf("string not expected %v", v)) // want "avoid use of panic"
	case int:
		panic("int not expected") // want "avoid use of panic"
	default:
		panic("unknown type") // want "avoid use of panic"
	}
}

func AnotherTestFunction() {
	unexportedTestFunctionWithPanic()
}

func ExportedTestFunctionWithoutPanic() interface{} {
	return unExportedTestFunctionWithoutPanic(10)
}

func unExportedTestFunctionWithoutPanic(a int) interface{} {
	return a
}
