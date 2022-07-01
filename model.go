package godsk

import "os"

type DSKFileFormat struct {
	fdesc    *os.File
	TRKS     [35][]byte
	Metadata map[string]string

	Version       int
	physicalTrack float32
	dataTrack     byte
	byteStreamPos  uint32
	revolution    int
}
