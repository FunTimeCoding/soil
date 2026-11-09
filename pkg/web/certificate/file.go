package certificate

import "github.com/funtimecoding/soil/pkg/web/constant"

func File() string {
	return resolve(
		constant.CertificatePathEnvironment,
		constant.CertificateFileName,
	)
}
