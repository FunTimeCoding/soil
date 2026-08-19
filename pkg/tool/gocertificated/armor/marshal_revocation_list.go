package armor

import (
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func MarshalRevocationList(b []byte) []byte {
	return pem.EncodeToMemory(
		&pem.Block{Type: constant.RevocationListBlock, Bytes: b},
	)
}
