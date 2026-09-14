package note

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewSlice(v []*gitlab.Note) []*Note {
	result := make([]*Note, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
