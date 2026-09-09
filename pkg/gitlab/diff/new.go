package diff

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.Diff) *Diff {
	return &Diff{
		OldPath: v.OldPath,
		NewPath: v.NewPath,
		Created: v.NewFile,
		Renamed: v.RenamedFile,
		Deleted: v.DeletedFile,
		Patch:   v.Diff,
		Raw:     v,
	}
}
