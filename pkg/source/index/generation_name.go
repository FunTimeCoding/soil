package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"runtime"
)

func generationName() string {
	sum := sha256.Sum256([]byte(formatVersion()))

	return join.Empty(runtime.Version(), "-", hex.EncodeToString(sum[:6]))
}
