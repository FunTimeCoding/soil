package constant

import "time"

const (
	ModelContextGreetTool     = "greet"
	ModelContextNameParameter = "name"
	ModelContextTestToken     = "alfa"
	ModelContextToolPrefix    = "mcp__"
	ModelContextToolSeparator = "__"
)

const (
	ParameterBody       = "body"
	ParameterIdentifier = "identifier"
	ParameterKey        = "key"
	ParameterLimit      = "limit"
	ParameterOffset     = "offset"
	ParameterMessage    = "message"
	ParameterName       = "name"
	ParameterQuery      = "query"
	ParameterRaw        = "raw"
	ParameterTitle      = "title"
	ParameterUnfiltered = "unfiltered"
)

const (
	MarkIdentifierKey = "identifier"

	MarkReadmeDocument = "docs://readme"

	MarkUserPrefix = "user://"

	MarkProfileSuffix = "/profile"
)

const (
	RecentRunsFail = "list recent runs"
	RunLookupFail  = "look up run"
)

const ModelContextAddress = ":10080"
const ModelContextSessionIdle = time.Hour
const NoSession = "no model context session on the request"
