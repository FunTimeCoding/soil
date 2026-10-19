package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gojellyfind",
	"Jellyfin media server bridge for search and playback",
	"gojellyfind",
).WithInstructions(
	"Jellyfin media server - search libraries, play music to browser sessions, control playback. Find a session with supports_remote: true before playing.",
)

const (
	ListLibraries   = "list_libraries"
	SearchItems     = "search_items"
	GetItem         = "get_item"
	GetEpisodes     = "get_episodes"
	GetTracks       = "get_tracks"
	ListSessions    = "list_sessions"
	Play            = "play"
	PlaybackCommand = "playback_command"
	SetVolume       = "set_volume"
)
