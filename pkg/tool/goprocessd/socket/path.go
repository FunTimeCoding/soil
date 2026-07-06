package socket

import (
	"crypto/sha256"
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func Path(procfilePath string) string {
	absolute := system.AbsolutePath(filepath.Dir(procfilePath))
	hash := sha256.Sum256([]byte(absolute))
	name := fmt.Sprintf("%x", hash[:8])

	return filepath.Join(
		directory(),
		fmt.Sprintf("%s.sock", name),
	)
}
