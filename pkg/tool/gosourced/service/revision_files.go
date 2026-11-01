package service

import "github.com/funtimecoding/soil/pkg/git"

func revisionFiles(
	directory string,
	revision string,
) (map[string][]byte, error) {
	if revision == "" {
		return git.IndexFiles(directory, isGoSource)
	}

	return git.RevisionFiles(directory, revision, isGoSource)
}
