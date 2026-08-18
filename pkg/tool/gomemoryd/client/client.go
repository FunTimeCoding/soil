package client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

type Client interface {
	VersionsSince(
		since string,
		limit int,
	) []VersionEntry
	SaveImpression(
		content string,
		source string,
	)
	Profile(topic string) string
	RedactedMemories() map[int64]bool
	Statistics() *client.Statistics
	Relations() []client.Relation
}
