package probe

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func Parse(output string) (*client.SystemSnapshot, error) {
	s := sections(output)
	lines, okay := s[constant.ReleaseSection]

	if !okay {
		return nil, fmt.Errorf("no %s section", constant.ReleaseSection)
	}

	r, e := release(lines)

	if e != nil {
		return nil, e
	}

	result := &client.SystemSnapshot{
		Release: r,
		Packages: installed(
			s[constant.PackagesSection],
			candidates(s[constant.UpgradableSection]),
		),
	}

	if k := s[constant.KernelSection]; len(k) > 0 {
		result.Kernel = &k[0]
	}

	result.ListsUpdated = lists(s[constant.ListsSection])

	if o := offers(s[constant.PolicySection]); len(o) > 0 {
		result.KernelOffers = &o
	}

	result.KernelMetapackageInstalled = metapackageInstalled(
		s[constant.PolicySection],
	)

	return result, nil
}
