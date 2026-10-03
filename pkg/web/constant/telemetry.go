package constant

// Reference: https://opentelemetry.io/docs/specs/semconv/http/http-spans/
const (
	TelemetryRequestMethod = "http.request.method"
	TelemetryPath          = "url.path"
	TelemetryScheme        = "url.scheme"

	TelemetryQuery  = "url.query"
	TelemetryRoute  = "http.route"
	TelemetryStatus = "http.response.status_code"

	TelemetryClient    = "client.address"
	TelemetryPeer      = "network.peer.address"
	TelemetryProtocol  = "network.protocol.version"
	TelemetryServer    = "server.address"
	TelemetryUserAgent = "user_agent.original"

	TelemetryBodySize     = "http.request.body.size"
	TelemetryBody         = "http.request.body"
	TelemetryHeaderPrefix = "http.request.header"
)
