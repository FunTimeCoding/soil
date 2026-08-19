package publish

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"path/filepath"
	"strings"
)

func secretSkeleton(path string) string {
	name := strings.TrimSuffix(filepath.Base(path), constant.SecretExtension)

	return fmt.Sprintf(
		"---\napiVersion: v1\nkind: Secret\nmetadata: {name: %s}\ntype: %s\n# noinspection SpellCheckingInspection\ndata:\n",
		name,
		constant.SecretType,
	)
}
