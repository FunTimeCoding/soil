package discussion

import (
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func New(v *gitlab.Discussion) *Discussion {
	return &Discussion{
		Identifier: v.ID,
		Individual: v.IndividualNote,
		Notes:      note.NewSlice(v.Notes),
	}
}
