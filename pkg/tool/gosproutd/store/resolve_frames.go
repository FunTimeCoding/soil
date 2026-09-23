package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"
)

func (s *Store) resolveFrames(
	session string,
	names []string,
) []*frame.Frame {
	var result []*frame.Frame

	for _, v := range names {
		f := frame.New(session, v)
		errors.PanicOnError(
			s.mapper.Where(
				"session = ? AND name = ?",
				session,
				v,
			).FirstOrCreate(f).Error,
		)
		result = append(result, f)
	}

	return result
}
