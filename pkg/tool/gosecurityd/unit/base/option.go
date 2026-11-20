package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/option"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"path/filepath"
)

func (s *Server) Option() *option.Security {
	s.t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(s.t.TempDir(), "security.sqlite")
	result.CacheDirectory = s.t.TempDir()
	result.NationalDatabase = func() face.NationalDatabase {
		return s.NationalDatabase
	}
	result.SecurityTracker = func() face.SecurityTracker {
		return s.SecurityTracker
	}
	result.ExtendedTracker = func() face.SecurityTracker {
		return s.ExtendedTracker
	}

	return result
}
