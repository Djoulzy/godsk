package main

import (
	"github.com/Djoulzy/godsk"
)

func main() {
	disk, err := godsk.InitContainer("Dos33.dsk")
	if err != nil {
		panic(err)
	}
	disk.Dump(true)
}
