package file

import "gitlab.com/gitlab-org/api/client-go/v2"

type File struct {
	Name       string
	Path       string
	Size       int64
	Encoding   string
	Content    string
	Reference  string
	Hash       string
	CommitHash string
	Raw        *gitlab.File
}
