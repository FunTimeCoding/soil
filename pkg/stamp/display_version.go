package stamp

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
	"strings"
)

func (s *Stamp) DisplayVersion() string {
	v := strings.TrimSuffix(s.Version, semver.Build(s.Version))

	if !module.IsPseudoVersion(v) {
		return v
	}

	if t := s.Tag(); t != "" {
		return fmt.Sprintf(constant.UntaggedFormat, t)
	}

	return constant.Untagged
}
