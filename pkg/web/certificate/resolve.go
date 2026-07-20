package certificate

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"path/filepath"
)

func resolve(
	override string,
	name string,
) string {
	result := environment.Fallback(
		override,
		filepath.Join(
			system.StorageDirectory(
				constant.CertificateName,
				false,
			),
			name,
		),
	)

	if !system.FileExists(result) {
		panic(join.Empty("certificate file missing: ", result))
	}

	return result
}
