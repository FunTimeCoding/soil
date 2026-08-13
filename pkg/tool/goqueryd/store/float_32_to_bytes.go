package store

import (
	"encoding/binary"
	"math"
)

func float32ToBytes(v []float32) []byte {
	b := make([]byte, len(v)*4)

	for i, f := range v {
		binary.LittleEndian.PutUint32(b[i*4:], math.Float32bits(f))
	}

	return b
}
