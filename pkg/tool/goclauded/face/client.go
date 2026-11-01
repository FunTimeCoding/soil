package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
)

type Client interface {
	Address() string
	ClaimLabel(sessionIdentifier string, key string, value string, from string) (bool, error)
	Close()
	DeleteSession(x context.Context, callsign string) error
	DeleteSessionByIdentifier(x context.Context, identifier string, params *client.DeleteSessionByIdentifierParams) (*client.DeleteReceiptResponse, error)
	GetChannel(x context.Context, params *client.GetChannelParams) (*client.CheckResponse, error)
	GetChannelCallsign(x context.Context, params *client.GetChannelCallsignParams) (*client.ChannelCallsignResponse, bool, error)
	GetCheck(x context.Context, params *client.GetCheckParams) (*client.CheckResponse, error)
	GetCost(x context.Context, params *client.GetCostParams) (*client.CostResponse, error)
	GetCoverage(x context.Context) (*client.CoverageResponse, error)
	GetMessages(x context.Context, params *client.GetMessagesParams) (*client.MessageReadResponse, error)
	GetResolve(x context.Context, params *client.GetResolveParams) (*client.ResolveResponse, []client.ResolveMatch, error)
	GetSessionCruise(x context.Context, identifier string) (*client.CruiseResponse, error)
	GetSessionDetail(x context.Context, identifier string) (*client.SessionDetailResponse, error)
	GetSessionLabel(x context.Context, identifier string) (*client.SessionLabelResponse, error)
	GetSessionMessages(x context.Context, identifier string) (*client.MessagesResponse, error)
	GetSessionPeek(x context.Context, identifier string) (*client.PeekResponse, error)
	GetSessionToolContext(x context.Context, identifier string, params *client.GetSessionToolContextParams) (*client.ToolContextResponse, error)
	GetSessionTools(x context.Context, identifier string) (*client.ToolsResponse, error)
	GetSessionWindow(x context.Context, identifier string, params *client.GetSessionWindowParams) (*client.WindowResponse, error)
	GetSessions(x context.Context, params *client.GetSessionsParams) (*client.SessionListResponse, error)
	GetSessionsBashDump(x context.Context) (*client.BashDumpResponse, error)
	GetSessionsFind(x context.Context, params *client.GetSessionsFindParams) (*client.FindResponse, error)
	GetSessionsHeatmap(x context.Context, params *client.GetSessionsHeatmapParams) (*client.HeatmapResponse, error)
	GetSessionsSearch(x context.Context, params *client.GetSessionsSearchParams) (*client.SearchResponse, error)
	GetStatus(x context.Context) (*client.StatusResponse, error)
	GetTimeline(x context.Context, params *client.GetTimelineParams) ([]client.TimelineEntry, error)
	GetUsage(x context.Context) (*client.UsageResponse, bool, error)
	Notify(callsign string, source string, body string) error
	PostAnnounce(x context.Context, body client.PostAnnounceJSONRequestBody) error
	PostBackfill(x context.Context, params *client.PostBackfillParams) (*client.BackfillResponse, error)
	PostEditSession(x context.Context, body client.PostEditSessionJSONRequestBody) error
	PostNotify(x context.Context, body client.PostNotifyJSONRequestBody) (*client.DeliveryResponse, error)
	PostRegister(x context.Context, body client.PostRegisterJSONRequestBody) (*client.RegisterResponse, error)
	PostRelease(x context.Context, body client.PostReleaseJSONRequestBody) error
	PostSend(x context.Context, body client.PostSendJSONRequestBody) (*client.DeliveryResponse, error)
	PostSessionContext(x context.Context, identifier string, body client.PostSessionContextJSONRequestBody) error
	PostSessionCruise(x context.Context, identifier string, body client.PostSessionCruiseJSONRequestBody) (*client.CruiseResponse, error)
	PostSessionEnd(x context.Context, body client.PostSessionEndJSONRequestBody) error
	PostSessionExport(x context.Context, identifier string) (*client.ExportResponse, error)
	PostSessionLabel(x context.Context, identifier string, body client.PostSessionLabelJSONRequestBody) (*client.LabelResponse, error)
	PostSessionPulse(x context.Context, identifier string, body client.PostSessionPulseJSONRequestBody) (*client.DeliveryResponse, error)
	PostSessionsExport(x context.Context) (*client.ExportResponse, error)
	PostSweep(x context.Context) (*client.SweepResponse, error)
	PostTurnEnd(x context.Context, body client.PostTurnEndJSONRequestBody) error
	Pulse(sessionIdentifier string, body string) (bool, error)
	PulseCruise(sessionIdentifier string, body string) (bool, error)
	PulseImmediate(sessionIdentifier string, body string) (bool, error)
	Reachable(x context.Context) error
	RecentSessions(limit int) ([]*target.Target, error)
	ReleaseLabel(sessionIdentifier string, key string, value string, from string) (bool, error)
	RemoveLabel(sessionIdentifier string, key string, from string) error
	SetLabel(sessionIdentifier string, key string, value string, from string) error
	Subscribe(name string, kinds []string, l *logger.Logger, r face.Reporter) (<-chan *stream_event.Event, func())
}
