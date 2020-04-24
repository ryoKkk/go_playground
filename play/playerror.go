package play

import (
	"fmt"

	"github.com/pkg/errors"
)

var errorTest = errors.New("This is a test")
var errorWrapped = errors.Wrap(errorTest, "wrapper here"+"abcd")
var errorStack = errors.WithStack(errorWrapped)

func ErrorEquality() {
	// newError := errors.New("This is a test")
	fmt.Printf("error equality: %v\n", errors.Is(errorWrapped, errorTest))
	fmt.Println("wrapped error: ", errorWrapped)
	fmt.Printf("error stack: %v\n", errorStack)
}

func ErrorWithMessage() {
	errorNewTest := errors.WithMessage(errorTest, "anotherTest")
	fmt.Println("before: ", errorTest)
	fmt.Println("with new message: ", errorNewTest)
}
