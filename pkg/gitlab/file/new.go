package file

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.File) *File {
	return &File{
		Name:       v.FileName,
		Path:       v.FilePath,
		Size:       v.Size,
		Encoding:   v.Encoding,
		Content:    v.Content,
		Reference:  v.Ref,
		Hash:       v.SHA256,
		CommitHash: v.CommitID,
		Raw:        v,
	}
}
