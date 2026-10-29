package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/signature"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) Upload(
	version string,
	repository string,
	architecture string,
	file string,
	body io.Reader,
) error {
	for _, name := range []string{version, repository, architecture, file} {
		if !bareName(name) {
			return validation.New("not a bare name: %s", name)
		}
	}

	if !strings.HasSuffix(file, ".apk") {
		return validation.New("must end with .apk: %s", file)
	}

	key, e := s.privateKey()

	if e != nil {
		return e
	}

	directory := filepath.Join(s.packageRoot, version, repository, architecture)

	if f := os.MkdirAll(directory, 0o755); f != nil {
		return f
	}

	staging, g := os.CreateTemp(directory, ".upload-*")

	if g != nil {
		return g
	}

	defer discard(staging.Name())
	_, h := io.Copy(staging, body)
	errors.LogClose(staging)

	if h != nil {
		return h
	}

	if i := signature.SignPackageWithKey(
		staging.Name(),
		key,
		s.signatureKey,
	); i != nil {
		return i
	}

	if j := os.Rename(staging.Name(), filepath.Join(directory, file)); j != nil {
		return j
	}

	return s.indexer.Index(directory, s.keyPath())
}
