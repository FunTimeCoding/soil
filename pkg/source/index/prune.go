package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/generation"
	"path"
	"path/filepath"
	"time"
)

func Prune(directory string) []string {
	current := generationName()
	result := generation.Prune(directory, current, ownedEntry)

	for _, build := range generation.Collect(
		filepath.Join(directory, current, constant.IndexBuildDirectory),
		buildName(extractorFingerprint()),
		time.Now(),
		constant.GenerationRefresh,
		ownedEntry,
	) {
		result = append(
			result,
			path.Join(current, constant.IndexBuildDirectory, build),
		)
	}

	return result
}
