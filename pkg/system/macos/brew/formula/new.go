package formula

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/response"
)

func New(r *response.Formula) *Formula {
	return &Formula{
		MonitorIdentifier: r.Name,
		Name:              r.Name,
		InstalledVersions: r.InstalledVersions,
		CurrentVersion:    r.CurrentVersion,
		Link: fmt.Sprintf(
			"https://formulae.brew.sh/formula/%s",
			r.Name,
		),
		Pinned: r.Pinned,
	}
}
