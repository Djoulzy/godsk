package godsk

import (
	"os"
)

func InitDskFile(fileName string) (*DSKFileFormat, error) {
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
	D.fdesc = f

	for i := 0; i < 35; i++ {
		dataStart = uint32(i) * 4096
		f.Seek(int64(dataStart), 0)

		D.TRKS[i] = make([]byte, 4096)
		f.Read(D.TRKS[i])
	}

	D.dataTrack = 0
	D.physicalTrack = 0
	D.bitStreamPos = 0
	D.revolution = 0
}

func (D *DSKFileFormat) Dump(full bool) {

}