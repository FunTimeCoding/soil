package component

import (
	"fmt"
	"hash/crc32"
)

func RowMark(
	prefix string,
	name string,
) string {
	return fmt.Sprintf("%s%08x", prefix, crc32.ChecksumIEEE([]byte(name)))
}
