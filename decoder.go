package godsk

var diskbyte [0x40]byte = [0x40]byte{0x96, 0x97, 0x9A, 0x9B, 0x9D, 0x9E, 0x9F, 0xA6, 0xA7, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF, 0xCB, 0xCD, 0xCE, 0xCF, 0xD3, 0xD6, 0xD7, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF, 0xE5, 0xE6, 0xE7, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF}

func (D *DSKFileFormat) encode6_2(dest, src []byte) {
	var c int
	var bit_reverse = []byte{0, 2, 1, 3}

	for c = 0; c < 84; c++ {
		dest[c] =
			bit_reverse[src[c]&3] | (bit_reverse[src[c+86]&3] << 2) | (bit_reverse[src[c+172]&3] << 4)
	}

	dest[84] = (bit_reverse[src[84]&3] << 0) | (bit_reverse[src[170]&3] << 2)
	dest[85] = (bit_reverse[src[85]&3] << 0) | (bit_reverse[src[171]&3] << 2)

	for c = 0; c < 256; c++ {
		dest[86+c] = src[c] >> 2
	}

	// Exclusive OR each byte with the one before it.
	dest[342] = dest[341]
	location := 342
	for {
		if location <= 1 {
			break
		}
		location--
		dest[location] ^= dest[location-1]
	}

	// Map six-bit values up to full bytes.
	for c = 0; c < 343; c++ {
		dest[c] = diskbyte[dest[c]]
	}
}

func (D *DSKFileFormat) code44_A(val byte) byte {
	return ((val >> 1) & 0x55) | 0xAA
}

func (D *DSKFileFormat) code44_B(val byte) byte {
	return (val & 0x55) | 0xAA
}

func (D *DSKFileFormat) pushByte(track_number byte, value byte) {
	D.TRKS[track_number][D.byteStreamPos] = value
	D.byteStreamPos++
}

func (D *DSKFileFormat) deserialise_track(src []byte, track_number byte, is_prodos bool) {
	var sector byte

	D.byteStreamPos = 0
	D.TRKS[track_number] = make([]byte, 6646)
	// Write gap one, which contains 48 self-sync bytes
	for loop := 0; loop < 48; loop++ {
		D.pushByte(track_number, 0xFF)
	}

	for sector = 0; sector < 16; sector++ {
		// Write the address field, which contains:
		//   - PROLOGUE (D5AA96)
		//   - VOLUME NUMBER ("4 AND 4" ENCODED)
		//   - TRACK NUMBER ("4 AND 4" ENCODED)
		//   - SECTOR NUMBER ("4 AND 4" ENCODED)
		//   - CHECKSUM ("4 AND 4" ENCODED)
		//   - EPILOGUE (DEAAEB)
		D.TMAP[track_number].Sectors[sector].StartByte = D.byteStreamPos

		D.pushByte(track_number, 0xD5)
		D.pushByte(track_number, 0xAA)
		D.pushByte(track_number, 0x96)

		D.pushByte(track_number, D.code44_A(0xFE))
		D.pushByte(track_number, D.code44_B(0xFE))

		D.pushByte(track_number, D.code44_A(track_number))
		D.pushByte(track_number, D.code44_B(track_number))
		D.pushByte(track_number, D.code44_A(sector))
		D.pushByte(track_number, D.code44_B(sector))
		D.pushByte(track_number, D.code44_A(0xFE^(track_number^sector)))
		D.pushByte(track_number, D.code44_B(0xFE^(track_number^sector)))

		D.pushByte(track_number, 0xDE)
		D.pushByte(track_number, 0xAA)
		D.pushByte(track_number, 0xEB)

		// Write gap two, which contains six self-sync bytes
		for loop := 0; loop < 6; loop++ {
			D.pushByte(track_number, 0xFF)
		}

		// Write the data field, which contains:
		//   - PROLOGUE (D5AAAD)
		//   - 343 6-BIT BYTES OF NIBBLIZED DATA, INCLUDING A 6-BIT CHECKSUM
		//   - EPILOGUE (DEAAEB)
		D.pushByte(track_number, 0xD5)
		D.pushByte(track_number, 0xAA)
		D.pushByte(track_number, 0xAD)

		var logical_sector byte
		if sector == 15 {
			logical_sector = 15
		} else {
			if is_prodos {
				logical_sector = (sector * 8) % 15
			} else {
				logical_sector = (sector * 7) % 15
			}
		}

		var contents = make([]byte, 343)
		D.encode6_2(contents, src[int(logical_sector)*256:])
		for _, data := range contents {
			D.pushByte(track_number, data)
		}

		D.pushByte(track_number, 0xDE)
		D.pushByte(track_number, 0xAA)
		D.pushByte(track_number, 0xEB)

		// Write gap three, which contains 27 self-sync bytes
		for loop := 0; loop < 27; loop++ {
			D.pushByte(track_number, 0xFF)
		}
		D.TMAP[track_number].Sectors[sector].ByteCount = D.byteStreamPos - D.TMAP[track_number].Sectors[sector].StartByte
	}
}
