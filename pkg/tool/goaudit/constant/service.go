package constant

const (
	DaemonFlag               = "daemon"
	MissingDaemonKey         = "missing_daemon"
	MissingDaemonText        = "lifecycle daemon without daemon/ - see service-tool.md, The daemon"
	MissingServiceKey        = "missing_service"
	MissingServiceText       = "daemon/ without service/ - see service-tool.md, Services"
	MixedServiceKey          = "mixed_service"
	MixedServiceText         = "service/ holds code beside the services in %s - one set of use cases or several, never both; see service-tool.md, Package Structure"
	RunBuildsClientKey       = "run_builds_client"
	RunBuildsClientText      = "run.go builds %s - the daemon builds it from a factory on the option"
	RunBuildsClockKey        = "run_builds_clock"
	RunBuildsClockText       = "run.go builds %s - the daemon builds its clocks and hands them out through Workers()"
	DaemonNotClosedKey       = "daemon_not_closed"
	DaemonNotClosedText      = "run.go builds the daemon but never defers its Close"
	SurfaceHoldsDaemonKey    = "surface_holds_daemon"
	SurfaceHoldsDaemonText   = "%s imports the daemon - a surface gets the services it calls; see service-tool.md, Services"
	ServiceImportsMapperKey  = "service_imports_mapper"
	ServiceImportsMapperText = "%s imports gorm - a transaction belongs to the store's Atomically; see service-tool.md, Stores"
	DaemonDirectory          = "daemon"
	DaemonPackage            = "daemon"
	DaemonCloseFile          = "close.go"
	ServiceDirectory         = "service"
	ServiceType              = "Service"
	WorkerPackage            = "worker"
	WorkerDirectory          = "worker"
	MapperImport             = "gorm.io/gorm"
	LifecyclePackage         = "lifecycle"
	CloseMethod              = "Close"
	OpenPrefix               = "Open"
)

var (
	RunPlumbing = []string{
		"daemon",
		"guard",
		"lifecycle",
		"logger",
		"metric",
		"option",
		"server",
		"ticker",
		"web",
	}
	Surfaces = []string{"server", "model_context", "web"}
)
