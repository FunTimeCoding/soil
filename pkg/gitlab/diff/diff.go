package diff

import "gitlab.com/gitlab-org/api/client-go/v3"

type Diff struct {
	OldPath string
	NewPath string
	Created bool
	Renamed bool
	Deleted bool
	Patch   string
	Raw     *gitlab.Diff // nil when built from a merge request diff
}
