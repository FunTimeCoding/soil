package client

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
)

type Client interface {
	VersionsSince(
		since string,
		limit int,
	) []version_entry.Entry
	SaveImpression(
		content string,
		source string,
	)
	Profile(topic string) string
	RedactedMemories() map[int64]bool
	Statistics() *client.Statistics
	Relations() []client.Relation
}
