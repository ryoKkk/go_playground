package play

import "fmt"

func init() {
	fmt.Println("-- init func in playmap")
}

func InitMap() {
	m := map[int]string{
		1: "Acky",
		2: "Becky",
	}
	fmt.Println("init map : ", m)
}

func EmptyKeyOfMap() {
	m := make(map[int]string)
	v, ok := m[2]
	fmt.Printf("existed : %v, empty key: %v\n", ok, v)
}

func RemoveKeyOfMap() {
	m := map[int]string{
		1: "Acky",
		2: "Becky",
	}
	fmt.Println("map: ", m)
	delete(m, 1)
	fmt.Println("Key removed: ", m)
}

func CopyAndSetMap() {
	m := map[int]string{
		1: "Acky",
		2: "Becky",
	}
	m2 := m
	m2[3] = "Cindy"
	m2[1] = "Ava"
	fmt.Println("m: ", m)
	fmt.Println("m2: ", m2)
}
