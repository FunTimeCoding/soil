package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func externalKey(
	path string,
	version string,
) string {
	sum := sha256.Sum256([]byte(join.Space(formatVersion(), path, version)))

	return hex.EncodeToString(sum[:])
}
