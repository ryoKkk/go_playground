package play

import (
	"fmt"
	"time"
)

func PlayDate() {
	t := time.Now()
	fmt.Printf("year:, %d, month: %d, dat: %d\n", t.Year(), t.Month(), t.Day())
}
