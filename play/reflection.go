package play

import (
	"fmt"
	"reflect"
)

func RefTypeOf() {
	r := reftype{}
	fmt.Printf("type: %v\n", reflect.TypeOf(r))
}

type reftype struct {
}
