package constant

const (
	LokiHostEnvironment      = "LOKI_HOST"
	LokiUserEnvironment      = "LOKI_USER"
	LokiPasswordEnvironment  = "LOKI_PASSWORD"
	LokiNamespaceEnvironment = "LOKI_NAMESPACE"
	LokiExcludeEnvironment   = "LOKI_EXCLUDE"
)

const (
	Success = "success"

	LokiBase       = "/loki/api/v1"
	LokiLabel      = "/label"
	LokiLabels     = "/labels"
	LokiPush       = "/push"
	LokiQuery      = "/query"
	LokiQueryRange = "/query_range"
	LokiSeries     = "/series"
	LokiStatistic  = "/index/stats"
	LokiValues     = "/values"
	// Stream
	Stdout           = "stdout"
	Stderr           = "stderr"
	LokiMaximumLimit = 5000
)

const (
	LokiTextType     = "text"
	LokiNotationType = "json"
	SlogMessage      = "msg"
)
