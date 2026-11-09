package certificate

import "github.com/funtimecoding/soil/pkg/web/constant"

func Key() string {
	return resolve(
		constant.CertificateKeyPathEnvironment,
		constant.CertificateKeyName,
	)
}
