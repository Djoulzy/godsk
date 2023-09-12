package godsk

import "os"

type DSKSectorMap struct {
	StartByte uint32
	ByteCount uint32
}

type DSKTrkMap struct {
	ByteCount uint32
	Sectors   [16]DSKSectorMap
}

type DSKFileFormat struct {
	fdesc    *os.File
	TRKS     [35][]byte
	TMAP     [35]DSKTrkMap
	Metadata map[string]string

	Version       int
	physicalTrack float32
	dataTrack     byte
	byteStreamPos uint32
	revolution    int
	output        string
}
