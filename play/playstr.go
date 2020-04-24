package play

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("-- init func in playstr")
}

func StringTrim() {
	fmt.Printf("trim space: '%s'\n", strings.TrimSpace("   "))
	fmt.Printf("trim space: '%s'\n", strings.TrimSpace(" ab"))
	fmt.Printf("trim space: '%s'\n", strings.TrimSpace("ac "))
	fmt.Printf("trim space: '%s'\n", strings.TrimSpace(" aadd "))
}

func StringEmpty() {
	s := ""
	fmt.Printf("'%s' is empty : %v\n", s, len(s) == 0)
	s = "aa"
	fmt.Printf("'%s' is empty : %v\n", s, len(s) == 0)
}

func StringEqual() {
	s1 := "abc"
	s2 := "abc"
	fmt.Println("string equals : ", s1 == s2)
}

func Backtick() {
	fmt.Printf("backtick: %s\n", `"Of course,", he said.`)
}

func StringCast() {
	var s interface{} = "abc"
	r := -1
	r, ok := s.(int)
	fmt.Printf("is OK : '%v', value: '%v'\n", ok, r)
}

func SliceString() {
	s := "abcdefg"
	fmt.Println(s[:5], s[5:])
}
