package play

import "fmt"

func PlayReturn() {
	name, addr, age := test()
	fmt.Printf("name : %s, address: %s, age: %d\n", name, addr, age)
}

func test() (string, string, int) {
	return "acky", "tokyo", 10
}
