package unit

import "runtime/debug"

func buildInformation(
	version string,
	settings ...debug.BuildSetting,
) *debug.BuildInfo {
	return &debug.BuildInfo{
		Main: debug.Module{
			Path:    "github.com/funtimecoding/soil",
			Version: version,
		},
		Settings: settings,
	}
}
