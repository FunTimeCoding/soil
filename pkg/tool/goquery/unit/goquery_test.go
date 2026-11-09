package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	stringsJoin "github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goquery/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
	"testing"
)

func TestWriteWithoutResultPrintsNothing(t *testing.T) {
	for _, arguments := range writes() {
		c := newMock(t)
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() { output, failure, code = run(t, c, arguments...) },
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
	}
}

func TestFailedCommandExitsWithItsMessage(t *testing.T) {
	file := join.Join(t.TempDir(), "alfa.md")
	system.SaveFile(file, "# alfa\n\nbravo\n")

	for _, arguments := range append(
		writes(),
		[]string{"search", "bravo"},
		[]string{soil.List, "alfa"},
		[]string{"index"},
		[]string{"embed"},
		[]string{"context-list"},
		[]string{"list-tags"},
		[]string{"list-metadata", "alfa"},
		[]string{"status"},
		[]string{"oversize"},
		[]string{"failed"},
		[]string{"rechunk"},
		[]string{"chunk", file},
	) {
		c := newMock(t)
		c.Failure = errors.New("foxtrot")
		_, failure, code := run(t, c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
	}
}

func TestDeletingUnknownCollectionIsNotFound(t *testing.T) {
	_, failure, code := run(t, newMock(t), "delete-collection", "zulu")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "collection not found: zulu", failure)
}

func TestRemovingUnknownContextIsNotFound(t *testing.T) {
	_, failure, code := run(t, newMock(t), "context-remove", "alfa", "zulu")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "context not found: alfa zulu", failure)
}

func TestRechunkEmitsWhatItQueuedAndEmbedded(t *testing.T) {
	c := newMock(t)
	c.Queued = []string{"alfa/kilo.md"}
	c.EmbedResult = &client.EmbedResult{Documents: 1, Chunks: 3}
	output, _, code := run(t, c, "rechunk")
	assert.Integer(t, 0, code)
	var r result.Rechunk
	notation.MustDecode(output, &r, false)
	assert.Strings(t, []string{"alfa/kilo.md"}, r.Queued)
	assert.FatalNotNil(t, r.Embedded)
	assert.Integer(t, 3, r.Embedded.Chunks)
	assert.Integer(t, 1, c.Rechunked)
	assert.Integer(t, 1, c.Embedded)
}

func TestFailedRechunkEmbedsNothing(t *testing.T) {
	c := newMock(t)
	c.Failure = errors.New("foxtrot")
	_, _, code := run(t, c, "rechunk")
	assert.Integer(t, 1, code)
	assert.Integer(t, 0, c.Embedded)
}

func TestOversizeEmitsReportWithItsSummary(t *testing.T) {
	c := newMock(t)
	c.OversizeReport = &client.OversizeReport{
		Allowance:      512,
		EmbedAllowance: 2048,
		Files: []client.OversizeFile{
			{
				Collection: "alfa",
				Path:       "kilo.md",
				Chunks: []client.OversizeChunk{
					{Index: 1, Tokens: 600, EmbedTokens: 600},
					{Index: 2, Tokens: 3000, EmbedTokens: 3000},
				},
			},
		},
	}
	output, _, code := run(t, c, "oversize")
	assert.Integer(t, 1, code)
	var r result.Oversize
	notation.MustDecode(output, &r, false)
	assert.FatalNotNil(t, r.Summary)
	assert.Integer(t, 1, r.Summary.Files)
	assert.Integer(t, 2, r.Summary.Chunks)
	assert.Integer(t, 1, r.Summary.OverEmbedder)
}

func TestNothingOversizeExitsZero(t *testing.T) {
	output, _, code := run(t, newMock(t), "oversize")
	assert.Integer(t, 0, code)
	var r result.Oversize
	notation.MustDecode(output, &r, false)
	assert.FatalNotNil(t, r.Summary)
	assert.Integer(t, 0, r.Summary.Files)
}

func TestFailedEmitsRefusedCount(t *testing.T) {
	c := newMock(t)
	c.FailedReport = &client.FailedReport{
		Chunks: []client.FailedChunk{{Path: "kilo.md"}, {Path: "lima.md"}},
	}
	output, _, code := run(t, c, "failed")
	assert.Integer(t, 1, code)
	var r result.Failed
	notation.MustDecode(output, &r, false)
	assert.FatalNotNil(t, r.Summary)
	assert.Integer(t, 2, r.Summary.Refused)
}

func TestStatusEmitsStatus(t *testing.T) {
	output, _, code := run(t, newMock(t), "status")
	assert.Integer(t, 0, code)
	var s client.Status
	notation.MustDecode(output, &s, false)
	assert.Integer(t, 1, s.TotalDocuments)
}

func TestChunkPrintsPreviewAsDocument(t *testing.T) {
	c := newMock(t)
	c.Preview = &client.ChunkPreview{
		Allowance:      512,
		EmbedAllowance: 2048,
		Model:          "kilo",
		Window:         512,
		Chunks: []client.PreviewChunk{
			{Index: 1, FirstLine: 1, LastLine: 9, Tokens: 600, EmbedTokens: 600},
		},
	}
	file := join.Join(t.TempDir(), "alfa.md")
	system.SaveFile(file, "# alfa\n\nbravo\n")
	output, failure, code := run(t, c, "chunk", file)
	assert.Integer(t, 1, code)
	assert.String(t, "", failure)
	assert.StringContains(t, "windows\n", output)
	assert.StringContains(t, "  OVER by 88\n", output)
	assert.StringContains(
		t,
		"1 chunks, 1 over 512 tokens (kilo, window 512) or the embedder's 2048\n",
		output,
	)
}

func TestCutWordListGivesOneNotice(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	file := join.Join(t.TempDir(), "alfa.md")
	system.SaveFile(file, "# alfa\n\nbravo\n\n# charlie\n\ndelta\n")
	_, _, code := run(t, newMock(t), "chunk", "--snapshot", file)
	assert.Integer(t, 0, code)
	var words []string

	for i := range 25 {
		words = append(words, strings.Repeat("x", i+1))
	}

	added := strings.Join(words, " ")
	system.SaveFile(
		file,
		stringsJoin.Empty(
			"# alfa\n\nbravo ",
			added,
			"\n\n# charlie\n\ndelta ",
			added,
			"\n",
		),
	)
	output, failure, code := run(t, newMock(t), "chunk", "--compare", file)
	assert.Integer(t, 0, code)
	assert.StringContains(t, "compared with snapshot\n", output)
	assert.StringContains(t, " ...\n", output)
	assert.String(t, "word lists cut at 20 in 2 sections\n", failure)
}
