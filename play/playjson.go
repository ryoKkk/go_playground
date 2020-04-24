package play

import (
	"encoding/json"
	"fmt"
)

func StringToEmpty() {
	input := `{"name": "ellery", "age": 10}`
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("map: ", result)
}

func PointerToJson() {
	addr := "Tokyo"
	spec := "Small"
	h := house{
		&addr,
		&spec,
	}
	r, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	fmt.Println("result : ", string(r))
}

type house struct {
	Addr *string
	Spec *string
}
