package file_sum

import "time"

func New(
	size int64,
	modified time.Time,
	sum [32]byte,
) *Sum {
	return &Sum{Size: size, Modified: modified, Sum: sum}
}
