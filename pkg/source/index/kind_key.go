package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"reflect"
)

func kindKey(
	key string,
	k *kind.Kind,
) string {
	sum := sha256.Sum256(
		[]byte(join.Space(key, k.Name, Shape(reflect.TypeOf(k.Value())))),
	)

	return hex.EncodeToString(sum[:])
}
