package utils

import (
	"encoding/binary"
	"fmt"
)

func bytesToInt64(data []byte) int64 {
	return int64(binary.BigEndian.Uint64(data))
}

func ConvertBytesToInt64Arr(data []byte) ([]int64, error) {
	if len(data)%8 != 0 {
		return nil, fmt.Errorf("invalid byte slice length: must be a multiple of 8")
	}

	result := make([]int64, len(data)/8)
	for i := 0; i < len(data); i += 8 {
		result[i/8] = bytesToInt64(data[i : i+8])
	}
	return result, nil
}
