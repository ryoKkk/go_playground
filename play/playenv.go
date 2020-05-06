package play

import (
	"fmt"
	"os"
	"strconv"
)

func EnvVar() {
	os.Setenv("PRODUCER_SIZE", "10")
	os.Setenv("CONSUMER_SIZE", "f")
	ps, _ := strconv.Atoi(os.Getenv("PRODUCER_SIZE"))
	fmt.Println("producer size: ", ps)
	cs, _ := strconv.Atoi(os.Getenv("CONSUMER_SIZE"))
	fmt.Println("consumer size: ", cs)
}
