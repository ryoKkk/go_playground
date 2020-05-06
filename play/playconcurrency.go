package play

import (
	"fmt"
	"sync"
	"time"
)

func PlayChannel() {
	ch := make(chan int, 10)
	quit := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("fib num: ", <-ch)
		}
		quit <- 10
	}()
	fibonacci(ch, quit)
}

func fibonacci(in chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select {
		case in <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func PlaySelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func PlayFanIn() {
	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	ch2 := make(chan int, 5)
	go func() {
		for i := 10; i < 15; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	ch3 := make(chan int, 8)
	go func() {
		for i := 100; i < 108; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	ch := merge(ch1, ch2, ch3)
	for i := range ch {
		fmt.Println("element : ", i)
	}
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	output := func(ch <-chan int) {
		for i := range ch {
			out <- i
		}
		wg.Done()
	}
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func PlayCloseChannel() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 4; i++ {
		v, ok := <-c
		fmt.Println("first element: ", v, ", ok: ", ok)
	}
}
