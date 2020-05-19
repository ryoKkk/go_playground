package play

import (
	"crypto/md5"
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
	b := [md5.Size]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15}
	fmt.Println("cast: ", string(b[:]))
}

func SliceString() {
	s := "abcdefg"
	fmt.Println(s[:5], s[5:])
}

func StringCapitalize() {
	s := "catalogCode"
	fmt.Printf("catalog code: %s\n", strings.Title(s))
}
