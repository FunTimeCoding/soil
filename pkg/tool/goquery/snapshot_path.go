package goquery

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
)

func snapshotPath(
	file string,
	create bool,
) string {
	directory := join.Join(
		constant.Identity.StorageDirectory(create),
		constant.SnapshotDirectory,
	)

	if create {
		system.MakeDirectory(directory)
	}

	sum := sha256.Sum256([]byte(system.AbsolutePath(file)))

	return join.Join(directory, hex.EncodeToString(sum[:]))
}
