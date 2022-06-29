package godsk

func (D *DSKFileFormat) GetNextByte() byte {
	// fmt.Printf("-- [%c] T:%02.02f (%d) Pos:%d    \r", wheel[count], W.physicalTrack, W.dataTrack, W.bitStreamPos)
	// count++
	// if count >= len(wheel) {
	// 	count = 0
	// }
	result := D.TRKS[D.dataTrack][D.bitStreamPos]
	return result
}

func (D *DSKFileFormat) GoToTrack(num float32) {
	newDataTrack, ok := W.TMAP.Map[num]
	if !ok {
		panic("bad track")
	}

	W.revolution = 0

	if newDataTrack == W.dataTrack {
		W.physicalTrack = num
		return
	}

	W.physicalTrack = num
	W.dataTrack = newDataTrack
	if W.bitStreamPos > 3 {
		W.bitStreamPos -= 4
	}
	// fmt.Printf("Move to T:%02.02f (%d) at pos %d\n", W.physicalTrack, W.dataTrack, W.bitStreamPos)
}

func (D *DSKFileFormat) Seek(offset float32) {
	var maxTrack float32
	destTrack := W.physicalTrack + offset
	// fmt.Printf("Seek Track offset %.02f -> %d\n", offset, W.TMAP.Map[destTrack])

	if W.Version >= 2 {
		maxTrack = 40
	} else {
		maxTrack = 35
	}

	if destTrack < 0 {
		destTrack = 0
	} else if destTrack > maxTrack {
		destTrack = maxTrack
	}
	W.GoToTrack(destTrack)
}
