package godsk

import (
	"os"
)

func InitContainer(fileName string) (*DSKFileFormat, error) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	tmp := DSKFileFormat{}
	tmp.init(file)

	return &tmp, err
}

func (D *DSKFileFormat) init(f *os.File) {
	var dataStart uint32
	var trkBuff = make([]byte, 4096)
	var i byte
	D.fdesc = f

	for i = 0; i < 35; i++ {
		dataStart = uint32(i) * 4096
		f.Seek(int64(dataStart), 0)

		f.Read(trkBuff)
		D.deserialise_track(trkBuff, i, false)
	}

	D.dataTrack = 0
	D.physicalTrack = 0
	D.byteStreamPos = 0
	D.revolution = 0
	D.Metadata = make(map[string]string)
}

func (D *DSKFileFormat) Dump(full bool) {

}
