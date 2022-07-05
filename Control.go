package godsk

import "fmt"

var count int = 0
var wheel []byte = []byte{'-', '\\', '|', '/'}
var pickbit = []byte{128, 64, 32, 16, 8, 4, 2, 1}

func (D *DSKFileFormat) IsWriteProtected() bool {
	return true
}

func (D *DSKFileFormat) GetMeta() map[string]string {
	return D.Metadata
}

func (D *DSKFileFormat) GetNextByte() byte {
	var result byte

	result = D.TRKS[D.dataTrack][D.byteStreamPos]
	D.byteStreamPos++
	if D.byteStreamPos > 6645 {
		D.byteStreamPos = 0
		D.revolution++
	}

	fmt.Printf("-- [%c] T:%02.02f (%d) Rev: %02d Pos:%d    \r", wheel[count], D.physicalTrack, D.dataTrack, D.revolution, D.byteStreamPos)
	count++
	if count >= len(wheel) {
		count = 0
	}
	return result
}

func (D *DSKFileFormat) FindSectorStart(trackNum byte) uint32 {
	var s int
	for s = 0; s < 15 && D.TMAP[int(trackNum)].Sectors[s].StartByte < D.byteStreamPos; s++ {
	}
	s -= 1
	if s < 0 {
		s = 0
	}
	if s > 15 {
		s = 15
	}
	return D.TMAP[int(trackNum)].Sectors[s].StartByte
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
	D.byteStreamPos = D.FindSectorStart(newDataTrack)
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

func (D *DSKFileFormat) DumpTrack(track float32) {
	var val byte

	D.GoToTrack(track)
	D.byteStreamPos = 0
	for i := 1; i <= 6646; i++ {
		val = D.GetNextByte()
		fmt.Printf("%02X ", val)
		if i%32 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func (D *DSKFileFormat) DumpTrackRaw(track float32) {

}
