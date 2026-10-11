package packager

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
)

func (p *Packager) calculateArchiveHash(path string) string {
	f := system.Open(path)
	defer errors.PanicClose(f)
	h := sha256.New()
	system.Copy(f, h)

	return hex.EncodeToString(h.Sum(nil))
}
