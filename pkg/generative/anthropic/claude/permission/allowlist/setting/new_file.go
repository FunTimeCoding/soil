package setting

func NewFile(p *Permissions) *File {
	return &File{Permissions: p}
}
