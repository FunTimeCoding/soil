package service

import (
	"encoding/base64"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/types/saved_capture"
	"os"
	"path/filepath"
)

func (s *Service) Screenshot(identifier int) (*saved_capture.Capture, error) {
	v, e := s.client.Screenshot(identifier)

	if e != nil {
		return nil, e
	}

	b, f := base64.StdEncoding.DecodeString(v.Content)

	if f != nil {
		return nil, fmt.Errorf("screenshot decode: %w", f)
	}

	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("firefox_%d.png", v.Identifier),
	)

	if g := os.WriteFile(path, b, 0600); g != nil {
		return nil, fmt.Errorf("screenshot write: %w", g)
	}

	return saved_capture.New(path, v.Title, len(b)), nil
}
