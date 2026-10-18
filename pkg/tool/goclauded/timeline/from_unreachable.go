package timeline

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func FromUnreachable(
	reason error,
	at string,
) *Entry {
	return &Entry{
		Timestamp: at,
		Kind:      constant.MemoryUnreachableKind,
		Metadata:  map[string]string{constant.Description: reason.Error()},
	}
}
