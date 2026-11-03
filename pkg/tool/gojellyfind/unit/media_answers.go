package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func mediaAnswers(
	x context.Context,
	f face.Client,
) []string {
	var result []string
	answer := func(
		v any,
		e error,
	) {
		result = append(result, answerOf(v, e))
	}
	answer(f.GetItem(x, ""))
	answer(f.GetItem(x, "missing"))
	answer(f.GetItem(x, "movie-1"))
	answer(f.GetEpisodes(x, ""))
	answer(f.GetEpisodes(x, "series-1"))
	answer(f.GetTracks(x, ""))
	answer(f.GetTracks(x, "album-1"))
	answer(f.ListLibraries(x))
	answer(f.ListSessions(x))
	answer(f.SearchItems(x, &client.SearchItemsParams{Q: "alfa"}))
	answer(
		f.SearchItems(
			x,
			&client.SearchItemsParams{Types: new("Episode, Audio")},
		),
	)
	answer(
		f.SearchItems(
			x,
			&client.SearchItemsParams{Page: new(2), PerPage: new(1)},
		),
	)
	answer(
		nil,
		f.Play(x, "", client.PlayJSONRequestBody{ItemIds: []string{"a"}}),
	)
	answer(nil, f.Play(x, "session-1", client.PlayJSONRequestBody{}))
	answer(
		nil,
		f.Play(
			x,
			"session-1",
			client.PlayJSONRequestBody{ItemIds: []string{"movie-1"}},
		),
	)
	answer(
		nil,
		f.PlaybackCommand(
			x,
			"session-1",
			client.PlaybackCommandJSONRequestBody{},
		),
	)
	answer(
		nil,
		f.PlaybackCommand(
			x,
			"session-1",
			client.PlaybackCommandJSONRequestBody{Command: "Pause"},
		),
	)
	answer(
		nil,
		f.SetVolume(x, "session-1", client.SetVolumeJSONRequestBody{Level: 101}),
	)
	answer(
		nil,
		f.SetVolume(x, "session-1", client.SetVolumeJSONRequestBody{Level: 40}),
	)

	return result
}
