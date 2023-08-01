package main

import (
	"fmt"

	"github.com/abhisheknaths/goillsutrated/fanin"
	"github.com/abhisheknaths/goillustrated/utils"
)

func main() {
	stream1, a := utils.GetReader()
	stream2, b := utils.GetReader()
	stream3, c := utils.GetReader()
	ac := new(utils.AtomicCounter)
	muxStream := fanin.Merge(stream1, stream2, stream3)
	for val := range muxStream {
		fmt.Println(val)
		ac.Add()
	}
	fmt.Printf(`Total lines from reader %[1]d %[3]sTotal lines processed %[2]d %[4]s`, a+b+c, ac.GetCount(), "\n", "\n")
}
