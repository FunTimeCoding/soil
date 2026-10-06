package stamp

import (
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
	"strings"
)

func (s *Stamp) Tag() string {
	if s.Version == constant.DefaultVersion {
		return ""
	}

	v := strings.TrimSuffix(s.Version, semver.Build(s.Version))

	if !module.IsPseudoVersion(v) {
		return v
	}

	base, e := module.PseudoVersionBase(v)

	if e != nil {
		return ""
	}

	return base
}
