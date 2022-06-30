package godsk

import "fmt"

var count int = 0
var wheel []byte = []byte{'-', '\\', '|', '/'}

func (D *DSKFileFormat) IsWriteProtected() bool {
	return true
}

func (D *DSKFileFormat) GetMeta() map[string]string {
	return D.Metadata
}

func (D *DSKFileFormat) GetNextByte() byte {
	result := D.TRKS[D.dataTrack][D.bitStreamPos]
	D.bitStreamPos++

	if D.bitStreamPos > 4095 {
		D.bitStreamPos = 0
	}

	fmt.Printf("-- [%c] T:%02.02f (%d) Pos:%d    \r", wheel[count], D.physicalTrack, D.dataTrack, D.bitStreamPos)
	count++
	if count >= len(wheel) {
		count = 0
	}
	return result
}

func (D *DSKFileFormat) GoToTrack(num float32) {
	newDataTrack := byte(num)

	D.revolution = 0

	if newDataTrack == D.dataTrack {
		D.physicalTrack = num
		return
	}

	D.physicalTrack = num
	D.dataTrack = newDataTrack
	// fmt.Printf("Move to T:%02.02f (%d) at pos %d\n", W.physicalTrack, W.dataTrack, W.bitStreamPos)
}

func (D *DSKFileFormat) Seek(offset float32) {
	var maxTrack float32
	destTrack := D.physicalTrack + offset
	// fmt.Printf("Seek Track offset %.02f -> %d\n", offset, W.TMAP.Map[destTrack])

	maxTrack = 35

	if destTrack < 0 {
		destTrack = 0
	} else if destTrack > maxTrack {
		destTrack = maxTrack
	}
	D.GoToTrack(destTrack)
}
