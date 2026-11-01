package retargeted_document

import "github.com/funtimecoding/soil/pkg/lint/types/retarget_change"

func New(
	path string,
	original string,
	content string,
	changes []*retarget_change.Change,
) *Document {
	return &Document{
		Path:     path,
		Original: original,
		Content:  content,
		Changes:  changes,
	}
}
