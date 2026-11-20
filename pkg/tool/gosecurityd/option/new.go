package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"github.com/funtimecoding/soil/pkg/vulnerability/nvd"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func New() *Security {
	return &Security{
		CacheDirectory:    constant.Identity.StorageDirectory(false),
		StaleAfter:        constant.StaleAfter,
		TrackerStaleAfter: constant.TrackerStaleAfter,
		Clock:             time.Now,
		NationalDatabase: func() face.NationalDatabase {
			return nvd.NewEnvironment()
		},
		SecurityTracker: func() face.SecurityTracker {
			return debian.New()
		},
		ExtendedTracker: func() face.SecurityTracker {
			return debian.NewWithBase(
				locator.New(vulnerabilityConstant.FreexianHost).Base(
					vulnerabilityConstant.FreexianBase,
				),
			)
		},
	}
}
