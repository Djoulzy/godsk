package main

import (
	"github.com/Djoulzy/godsk"
)

func main() {
	disk, err := godsk.InitDskFile("Dos33.dsk")
	if err != nil {
		panic(err)
	}
	disk.Dump(true)
}
