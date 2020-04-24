package play

import "fmt"

var initialized = 10
var uninitialized int

// memo: imported packages initialized -> variable initialized -> init function
func init() {
	fmt.Printf("-- init func: initialized: %v\n", initialized)
	fmt.Printf("-- init func: uninitialized: %v\n", uninitialized)
}

func InitFunc() {
	fmt.Printf("initialized: %v\n", initialized)
	fmt.Printf("uninitialized: %v\n", uninitialized)
}
