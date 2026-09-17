package pointer

import "path"

func Ancestors(source string) []string {
	var result []string
	directory := path.Dir(path.Dir(source))

	for directory != "." && directory != "/" && directory != "" {
		result = append(result, directory)
		directory = path.Dir(directory)
	}

	return result
}
