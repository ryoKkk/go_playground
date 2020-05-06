package main

import "fmt"

func main() {
	/**
	fmt.Println("------- string ------")
	play.SliceString()
	fmt.Println("------- map ------")
	play.CopyAndSetMap()
	fmt.Println("------- struct ------")
	play.DefaultStructField()
	fmt.Println("------- package -------")
	play.InitFunc()
	fmt.Println("------- error -------")
	play.ErrorEquality()
	fmt.Println("------- attribute value -------")
	play.MapToAttrVal()
	play.StructToAttrVal()
	play.AttrValToMap()
	fmt.Println("------- time -------")
	play.DateToString()
	fmt.Println("------- slice -------")
	play.AppendSlice()
	fmt.Println("------- rand -------")
	play.PlayRand()
	fmt.Println("------- dynamodb -------")
	*/
	// play.DDBGetItem()
	// play.DDBUpdateItem()
	/**
	fmt.Println("------- json -------")
	play.StringToEmpty()
	play.PointerToJson()
	fmt.Println("------- reflect -------")
	play.RefTypeOf()
	fmt.Println("------- log -------")
	play.LogInfo()
	fmt.Println("------- interface -------")
	h := play.DefaultHandler{}
	(&h).OnSuccess()
	// NG
	// play.Succeeds(h)
	play.Succeeds(&h)
	*/
	p := receiver{
		"Ellery",
		20,
	}
	p.SayHello("How are you ?")
	p.Dummy()
	(receiver).SayHello(p, "yes")
}

type receiver struct {
	Name string
	Age  int
}

func (p receiver) SayHello(greeting string) {
	fmt.Println(greeting + " from " + p.Name)
}

func (p receiver) Dummy() {

}
