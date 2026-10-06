package index

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"strings"
)

func Shape(types ...reflect.Type) string {
	var b strings.Builder
	seen := make(map[reflect.Type]bool)

	for _, t := range types {
		writeShape(&b, t, seen)
	}

	result := sha256.Sum256([]byte(b.String()))

	return hex.EncodeToString(result[:])
}
