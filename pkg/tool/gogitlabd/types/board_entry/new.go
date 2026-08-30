package board_entry

import "time"

func New(
	project string,
	projectLink string,
	reference string,
	status string,
	identifier int64,
	link string,
	updated time.Time,
) *Entry {
	return &Entry{
		Project:     project,
		ProjectLink: projectLink,
		Reference:   reference,
		Status:      status,
		Identifier:  identifier,
		Link:        link,
		Updated:     updated,
	}
}
