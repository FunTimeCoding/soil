package backup

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"io"
	"os"
	"path/filepath"
)

func (s *Store) WriteChecksum(path string) (string, int64, error) {
	f, e := os.Open(path)

	if e != nil {
		return "", 0, e
	}

	defer errors.PanicClose(f)
	h := sha256.New()
	size, g := io.Copy(h, f)

	if g != nil {
		return "", 0, g
	}

	result := hex.EncodeToString(h.Sum(nil))

	if w := os.WriteFile(
		join.Empty(path, constant.ChecksumExtension),
		[]byte(fmt.Sprintf(constant.ChecksumLine, result, filepath.Base(path))),
		constant.BackupFileMode,
	); w != nil {
		return "", 0, w
	}

	return result, size, nil
}
