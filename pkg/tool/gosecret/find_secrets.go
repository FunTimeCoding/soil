package gosecret

import (
	"os"
	"path/filepath"
)

func FindSecrets(directory string) ([]string, error) {
	var secrets []string
	e := filepath.Walk(
		directory,
		func(
			path string,
			i os.FileInfo,
			f error,
		) error {
			if f != nil {
				return f
			}

			if i.IsDir() {
				return nil
			}

			if !isYAMLFile(path) {
				return nil
			}

			isSecret, e := IsSecretManifest(path)

			if e != nil {
				return e
			}

			if isSecret {
				secrets = append(secrets, path)
			}

			return nil
		},
	)

	return secrets, e
}
