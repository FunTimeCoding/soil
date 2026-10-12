package gosecret

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/face"
	"os"
	"path/filepath"
	"slices"
)

func Migrate(
	s face.SecretStore,
	manifest string,
	namespace string,
	store string,
) (*MigrateResult, error) {
	m, e := readManifest(manifest)

	if e != nil {
		return nil, e
	}

	values, e := decodePayload(m.Payload)

	if e != nil {
		return nil, e
	}

	if len(values) == 0 {
		return nil, validation.New(
			"Secret manifest holds no keys: %s",
			manifest,
		)
	}

	decoded := GetDecodedPath(manifest)
	inSync, e := decodedMatches(decoded, values)

	if e != nil {
		return nil, e
	}

	if !inSync {
		return nil, validation.New(
			"%s differs from its manifest - encode or decode first",
			decoded,
		)
	}

	if namespace == "" {
		namespace = filepath.Base(filepath.Dir(manifest))
	}

	result := &MigrateResult{
		Path: join.Empty(namespace, constant.PathSeparator, m.Metadata.Name),
	}
	result.Keys, e = s.Put(result.Path, values)

	if e != nil {
		return nil, e
	}

	for k, v := range values {
		stored, f := s.Reveal(result.Path, k)

		if f != nil {
			return nil, f
		}

		if stored != v {
			return nil, fmt.Errorf(
				"stored value differs from the manifest: %s",
				k,
			)
		}
	}

	for _, k := range result.Keys {
		if _, okay := values[k]; !okay {
			result.Extra = append(result.Extra, k)
		}
	}

	slices.Sort(result.Extra)
	system.SaveFile(
		manifest,
		ExternalSecret(m.Metadata.Name, result.Path, store, m.Type),
	)

	if system.FileExists(decoded) {
		if f := os.Remove(decoded); f != nil {
			return nil, fmt.Errorf("remove decoded file: %w", f)
		}

		result.Removed = decoded
	}

	return result, nil
}
