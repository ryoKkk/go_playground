package csp

import (
	"fmt"
)

func CspCopy() {
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
		fmt.Println("element : ", i)
		east <- i
	}
	for i := range east {
		fmt.Println("element in east: ", i)
	}
}

func CspSquash() {
	west := make(chan string, 10)
	go func() {
		west <- "a"
		west <- "b"
		west <- "c"
		west <- "*"
		west <- "*"
		west <- "d"
		west <- "*"
		west <- "e"
		west <- "*"
		west <- "f"
		west <- "*"
		close(west)
	}()
	east := make(chan string, 10)
	go func() {
		defer close(east)
		for {
			c, ok := <-west
			if !ok {
				break
			}
			if c == "*" {
				d, ok := <-west
				if !ok {
					east <- c
					break
				}
				if d == "*" {
					east <- "|"
				} else {
					east <- c
					east <- d
				}
			} else {
				east <- c
			}
		}
	}()
	for i := range east {
		fmt.Printf(i)
	}
	fmt.Println()
}

func CspDivision() {
	c := make(chan pair)
	division := func(ch <-chan pair) <-chan pair {
		r := make(chan pair)
		p := <-ch
		rem, divisor := p.x, p.y
		quo := 0
		go func() {
			for rem >= divisor {
				rem -= divisor
				quo++
			}
			r <- pair{x: rem, y: quo}
			close(r)
		}()
		return r
	}
	go func() {
		c <- pair{100, 3}
	}()
	r := division(c)
	for e := range r {
		fmt.Printf("remainder: %d, quotient: %d\n", e.x, e.y)
	}
}

type pair struct {
	x, y int
}

/**
func CsvFactorial() {
	user := make(chan int)
	factorial := func(input <-chan int) <-chan int {
		ch := make(chan int)

		return ch
	}
}
*/

func S42_Factorial(fac []chan int, limit int) {
	for i := 0; i < limit; i++ {
		go func(i int) {
			n := <-fac[i]
			if n == 0 {
				fac[i] <- 1
			} else if n > 0 {
				// Note that here we check if i equals limit.
				// The original solution in the paper fails to terminate
				// if user input is equal or higher than the given limit.
				if i == limit-1 {
					fac[i] <- n
				} else {
					fac[i+1] <- n - 1
					r := <-fac[i+1]
					fac[i] <- n * r
				}
			}
			close(fac[i])
		}(i)
	}
}

func PlayCspFactorial() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	S42_Factorial([]chan int{ch1, ch2, ch3, ch4}, 3)
	ch1 <- 4
	for e := range ch1 {
		fmt.Println(e)
	}
}
