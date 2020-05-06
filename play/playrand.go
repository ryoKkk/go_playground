package play

import (
	"math/rand"
	"fmt"
)

func PlayRand() {
	for i := 0; i < 1; i++ {
		fmt.Println("rand: ", rand.Intn(10))
	}
}
