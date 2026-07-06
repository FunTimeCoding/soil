package vocabulary

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/constant"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download() ([]byte, error) {
	r, e := http.Get(constant.VocabularyLink)

	if e != nil {
		return nil, fmt.Errorf("failed to download vocabulary: %w", e)
	}

	defer errors.PanicClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to download vocabulary: %s",
			r.Status,
		)
	}

	content, e := io.ReadAll(r.Body)

	if e != nil {
		return nil, e
	}

	path := cachePath()

	if path == "" {
		return content, nil
	}

	if f := os.MkdirAll(filepath.Dir(path), 0o755); f != nil {
		return content, nil
	}

	errors.PanicOnError(os.WriteFile(path, content, 0o644))

	return content, nil
}
