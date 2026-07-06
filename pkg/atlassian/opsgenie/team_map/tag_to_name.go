package team_map

import "github.com/funtimecoding/soil/pkg/face"

func (m *Map) TagToName(f face.SliceAlias) {
	m.tagToName = f
}
