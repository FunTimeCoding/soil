package index

import (
	"crypto/sha256"
	"encoding/hex"
)

func buildName(build string) string {
	sum := sha256.Sum256([]byte(build))

	return hex.EncodeToString(sum[:8])
}
