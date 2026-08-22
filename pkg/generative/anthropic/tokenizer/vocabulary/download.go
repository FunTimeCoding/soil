package vocabulary

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download() ([]byte, error) {
	r, e := http.Get(constant.AnthropicVocabularyLink)

	if e != nil {
		return nil, fmt.Errorf("download vocabulary: %w", e)
	}

	defer errors.PanicClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download vocabulary: %s", r.Status)
	}

	content, e := io.ReadAll(r.Body)

	if e != nil {
		return nil, e
	}

	path := cachePath()
	system.MakeDirectory(filepath.Dir(path))
	errors.PanicOnError(os.WriteFile(path, content, 0o644))

	return content, nil
}
