package play

import (
	"fmt"
	"time"
)

func DateToString() {
	n := time.Now()
	fmt.Printf("now: %v\n", n.String())
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().Format("20060102150405.000"))
	}

	t, err := time.Parse("20060102150405", "20200420163300")
	if err != nil {
		panic(err)
	}
	t = t.In(time.Local)
	fmt.Println("parsed timestamp: ", t)
}

func DateOperation() {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	fmt.Println("tomorrow: ", tomorrow)
	d := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	fmt.Println("equal: ", tomorrow.Equal(d))
}
