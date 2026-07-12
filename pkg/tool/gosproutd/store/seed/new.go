package seed

import "time"

func New(
	name string,
	path string,
	contentHash string,
	content string,
	position int,
	modifiedAt time.Time,
) *Seed {
	return &Seed{
		Name:        name,
		Path:        path,
		ContentHash: contentHash,
		Content:     content,
		Position:    position,
		ModifiedAt:  modifiedAt,
	}
}
