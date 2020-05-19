package play

import "fmt"

type Handler interface {
	OnSuccess()
}

type DefaultHandler struct {
}

func (h DefaultHandler) OnSuccess() {
	fmt.Println("On Success")
}

func Succeeds(h Handler) {
	h.OnSuccess()
}

func PlayMethodSugar() {
	h := DefaultHandler{}
	h.OnSuccess()
	(DefaultHandler).OnSuccess(h)
}
