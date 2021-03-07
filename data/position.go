package positions

func encodePositions(x int64, y int64, z int64) (uint64) {
	return uint64(((x & 0x3FFFFFF) << 38) | ((y & 0xFFF) << 26) | (z & 0x3FFFFFF))
}

func decodePositions(pos uint64) (int64, int64, int64) {
	return int64(pos >> 38), int64((pos >> 26) & 0xFFF), int64(pos << 38 >> 38)
}
