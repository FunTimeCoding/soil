package scan_tester

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan/audit_configuration"

func AlfaConfiguration() *audit_configuration.Configuration {
	configuration := audit_configuration.New()
	configuration.ModelContext = map[string]string{"alfa": "pkg/tool/goalfad"}

	return configuration
}
