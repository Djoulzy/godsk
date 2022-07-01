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

	// src := disk.TRKS[0]
	// dest := godsk.deserialise_track(src, 0, false)

	// for _, data := range dest {
	// 	fmt.Printf("%02X ", data)
	// }
}
