package license

import "github.com/funtimecoding/soil/pkg/go_mod/constant"

func Family(identifiers []string) string {
	result := constant.LicenseUnknown

	for _, i := range identifiers {
		switch {
		case hasPrefix(i, constant.CopyleftLicenses):
			return constant.Copyleft
		case hasPrefix(i, constant.WeakCopyleftLicenses):
			result = constant.WeakCopyleft
		case hasPrefix(i, constant.PermissiveLicenses) &&
			result == constant.LicenseUnknown:
			result = constant.Permissive
		}
	}

	return result
}
