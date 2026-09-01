package entry_detail

import "time"

func New(
	identifier string,
	path string,
	title string,
	fields map[string]string,
	modifiedAt time.Time,
) *Detail {
	return &Detail{
		Identifier: identifier,
		Path:       path,
		Title:      title,
		Fields:     fields,
		ModifiedAt: modifiedAt,
	}
}
