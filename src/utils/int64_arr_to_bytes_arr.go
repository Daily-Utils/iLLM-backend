package utils

import "encoding/binary"

func int64ToBytes(num int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))
	return buf
}

func ConvertInt64ToBytesArr(arr []int64) []byte {
	contextBytes := make([]byte, len(arr)*8)
	for i, val := range arr {
		copy(contextBytes[i*8:(i+1)*8], int64ToBytes(val))
	}
	return contextBytes
}
