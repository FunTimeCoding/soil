package service_tester

import (
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func WriteLogCache(
	directory string,
	names ...string,
) {
	logs := map[string]any{}

	for _, n := range names {
		logs[n] = map[string]any{"FileName": n, "MapId": 38}
	}

	system.WriteFile(
		filepath.Join(directory, gw2Constant.LogFile),
		notation.Marshal(map[string]any{"Version": 1, "LogsByFilename": logs}),
		0o600,
	)
}
