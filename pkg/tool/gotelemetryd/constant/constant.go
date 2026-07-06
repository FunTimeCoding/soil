package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gotelemetryd",
	"Usage telemetry collector",
	"gotelemetryd",
).WithInstructions(
	"Usage telemetry across all surfaces (MCP, CLI, REST, web). Query tool usage heatmaps, filter by tool, surface, actor, and time window.",
)

const (
	Query   = "query"
	Summary = "summary"

	Tool    = "tool"
	Surface = "surface"
	Actor   = "actor"
	Kind    = "kind"
	Since   = "since"
	Until   = "until"
	Limit   = "limit"
	GroupBy = "group_by"

	SurfaceModelContext = "model_context"
	SurfaceCommandLine  = "command_line"
	SurfaceWeb          = "web"
	SurfaceWebService   = "web_service"

	OutcomeSuccess = "success"
	OutcomeError   = "error"

	QueryFailed   = "failed to query events"
	SummaryFailed = "failed to build summary"

	HeatmapTitle = "Heatmap"
	HeatmapPath  = "/"
	EventsTitle  = "Events"
	EventsPath   = "/events"
)
