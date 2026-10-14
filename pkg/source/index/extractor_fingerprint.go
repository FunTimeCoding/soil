package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
	"sync"
)

var extractorFingerprint = sync.OnceValue(
	func() string {
		path, e := os.Executable()

		if e != nil {
			return ""
		}

		f, e := os.Open(path)

		if e != nil {
			return ""
		}

		defer errors.LogClose(f)
		h := sha256.New()

		if _, e = io.Copy(h, f); e != nil {
			return ""
		}

		return hex.EncodeToString(h.Sum(nil))
	},
)
