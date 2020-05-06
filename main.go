package main

import (
	"fmt"

	"github.com/ryoKkk/go_playground/play"
)

func main() {
	fmt.Println("------- string ------")
	play.StringCast()
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
	fmt.Println("------- env var -------")
	play.EnvVar()
	fmt.Println("------- aws s3 -------")
	play.DownloadFromS3()
	play.PathTest()
	fmt.Println("------- file -------")
	play.PlayFileMode()
	fmt.Println("------- channel -------")
	//play.PlayChannel()
	//play.PlaySelect()
	//play.PlayFanIn()
	//play.PlayCloseChannel()
	play.PlayConcurrentMD5()
	fmt.Println("------- return -------")
	//play.PlayReturn()
	fmt.Println("------- dynamodb -------")
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
}
