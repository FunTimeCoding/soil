package gosecret

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
)

func ExternalSecret(
	name string,
	path string,
	store string,
	secretType string,
) string {
	template := ""

	if secretType != "" && secretType != constant.OpaqueType {
		template = fmt.Sprintf(constant.TemplateTypeFormat, secretType)
	}

	return fmt.Sprintf(
		constant.ExternalSecretFormat,
		name,
		path,
		store,
		template,
	)
}
