package restoration

func New(
	database string,
	path string,
	clean bool,
) *Restoration {
	return &Restoration{Database: database, Path: path, Clean: clean}
}
