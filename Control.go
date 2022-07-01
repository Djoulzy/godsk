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

// func (D *DSKFileFormat) getNextBit() byte {
// 	// Lecture d'un track vide
// 	// fmt.Printf("DataTrack: %v\n", W.dataTrack)

// 	// D.bitStreamPos = D.bitStreamPos % 50304

// 	targetByte := D.bitStreamPos >> 3
// 	targetBit := D.bitStreamPos & 7

// 	res := (D.TRKS[D.dataTrack][targetByte] & pickbit[targetBit]) >> (7 - targetBit)

// 	D.bitStreamPos++
// 	if D.bitStreamPos > 50304 {
// 		D.bitStreamPos = 0
// 		D.revolution++
// 	}
// 	return res
// 	// return 0
// }

func (D *DSKFileFormat) GetNextByte() byte {
	var result byte

	// result = 0
	// for bit = 0; bit == 0; bit = D.getNextBit() {
	// }
	// result = 0x80 // the bit we just retrieved is the high bit
	// for i := 6; i >= 0; i-- {
	// 	result |= D.getNextBit() << i
	// }

	result = D.TRKS[D.dataTrack][D.byteStreamPos]
	D.byteStreamPos++
	if D.byteStreamPos > 6645 {
		D.byteStreamPos = 0
	}

	fmt.Printf("-- [%c] T:%02.02f (%d) Pos:%d    \r", wheel[count], D.physicalTrack, D.dataTrack, D.byteStreamPos)
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
