package main

import (
	"fmt"

	"github.com/abhisheknaths/goillsutrated/fanin"
	"github.com/abhisheknaths/goillustrated/fanout"
	"github.com/abhisheknaths/goillustrated/utils"
)

func main() {
	stream1, a := utils.GetReader()
	stream2, b := utils.GetReader()
	stream3, c := utils.GetReader()

	unifiedStream := fanin.Merge(stream1, stream2, stream3)

	done, processedCounter := fanout.Split(unifiedStream, 5)
	<-done
	fmt.Println("All done")
	fmt.Printf(`Total lines read %d%s`, a+b+c, "\n")
	fmt.Printf(`Total lines processed %d%s`, processedCounter.GetCount(), "\n")
}
