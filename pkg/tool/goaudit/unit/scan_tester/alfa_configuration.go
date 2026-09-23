package scan_tester

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan"

func AlfaConfiguration() *scan.Configuration {
	configuration := scan.NewConfiguration()
	configuration.ModelContext = map[string]string{"alfa": "pkg/tool/goalfad"}

	return configuration
}
