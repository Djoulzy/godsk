package godsk

import "os"

type DSKFileFormat struct {
	fdesc *os.File
	TRKS  [35][]byte

	Version       int
	physicalTrack float32
	dataTrack     byte
	bitStreamPos  uint32
	revolution    int
}
