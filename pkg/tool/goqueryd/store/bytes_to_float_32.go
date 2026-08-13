package store

import (
	"encoding/binary"
	"math"
)

func bytesToFloat32(b []byte) []float32 {
	result := make([]float32, len(b)/4)

	for i := range result {
		result[i] = math.Float32frombits(binary.LittleEndian.Uint32(b[i*4:]))
	}

	return result
}
