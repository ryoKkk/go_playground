package play

import "fmt"

func DefaultStructField() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{}
	fmt.Println("(NO GOOD) default struct field: ", p)
}
