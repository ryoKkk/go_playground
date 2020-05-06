package csp

import "fmt"

func Func243() {
	len := 10
	west := make(chan rune, len)
	go func() {
		for i := 'a'; i < 'a'+10; i++ {
			west <- i
		}
		close(west)
	}()
	east := make(chan rune, len)
	for i := range west {
		east <- i
	}
	for i := range east {
		fmt.Println("element in east: ", i)
	}
}
