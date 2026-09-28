package digest

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func hexDigest(parts ...string) string {
	sum := md5.Sum([]byte(strings.Join(parts, ":")))

	return hex.EncodeToString(sum[:])
}
