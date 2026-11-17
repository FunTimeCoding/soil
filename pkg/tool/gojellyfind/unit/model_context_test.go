package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"testing"
)

func TestModelContext(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.Count(t, 9, c.ListTools())
	s.Source.AddLibrary(library.New("lib-1", "Music", "music"))
	assert.StringContains(
		t,
		"Music",
		c.MustCallTool(constant.ListLibraries, map[string]any{}),
	)
	i := item.Stub()
	i.Identifier = "item-1"
	i.Name = "Dark Side of the Moon"
	i.Type = "MusicAlbum"
	s.Source.AddItem(i)
	j := item.Stub()
	j.Identifier = "item-2"
	j.Name = "Breathe"
	j.Type = "Audio"
	s.Source.AddItem(j)
	assert.StringContains(
		t,
		"Dark Side",
		c.MustCallTool(constant.SearchItems, map[string]any{"q": "dark"}),
	)
	assert.StringContains(
		t,
		"item-1",
		c.MustCallTool(constant.GetItem, map[string]any{"id": "item-1"}),
	)
	assert.StringContains(
		t,
		"Breathe",
		c.MustCallTool(
			constant.GetTracks,
			map[string]any{"album_id": "item-1"},
		),
	)
	k := item.Stub()
	k.Identifier = "ep-1"
	k.Name = "Pilot"
	k.Type = "Episode"
	k.SeriesName = "Breaking Bad"
	k.SeasonNumber = 1
	k.EpisodeNumber = 1
	s.Source.AddItem(k)
	assert.StringContains(
		t,
		"Pilot",
		c.MustCallTool(
			constant.GetEpisodes,
			map[string]any{"series_id": "series-1"},
		),
	)
	l := session.Stub()
	l.Identifier = "sess-1"
	l.DeviceName = "Firefox"
	l.Client = "Jellyfin Web"
	l.SupportsRemote = true
	s.Source.AddSession(l)
	assert.StringContains(
		t,
		"Firefox",
		c.MustCallTool(constant.ListSessions, map[string]any{}),
	)
	assert.StringContains(
		t,
		"playback started",
		c.MustCallTool(
			constant.Play,
			map[string]any{
				"session_id": "sess-1",
				"item_ids":   []string{"item-1"},
			},
		),
	)
	assert.StringContains(
		t,
		"sent",
		c.MustCallTool(
			constant.PlaybackCommand,
			map[string]any{"session_id": "sess-1", "command": "Pause"},
		),
	)
	assert.StringContains(
		t,
		"volume set",
		c.MustCallTool(
			constant.SetVolume,
			map[string]any{"session_id": "sess-1", "level": 50},
		),
	)
}
