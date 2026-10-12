package import_entry

func New(
	path string,
	declared string,
	local string,
	module string,
	standard bool,
) *Entry {
	return &Entry{
		Path:     path,
		Declared: declared,
		Local:    local,
		Module:   module,
		Standard: standard,
		Stays:    true,
	}
}
