package play

import "fmt"

func AppendSlice() {
	a := make([]byte, 0, 32)
	a = append(a, 'b', 'a')
	a1 := append(a, 'd')
	fmt.Println(string(a1))
	a2 := append(a, 'g')

	fmt.Println(string(a1))
	fmt.Println(string(a2))
}
