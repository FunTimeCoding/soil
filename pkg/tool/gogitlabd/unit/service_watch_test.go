package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"strings"
	"testing"
	"time"
)

func TestWatchNeedsGoclauded(t *testing.T) {
	o := service_tester.New(t)
	o.SeedWatched()
	_, _, _, e := o.Service().SubscribeMergeRequest(
		t.Context(),
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
		"",
		goclaudedConstant.DeliveryImmediate,
	)
	assert.True(t, not_configured.Is(e))
}

func TestWatchBaselineAnnouncesNothing(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Alfa().SeedNote(7, 3, "delta", "earlier", false, time.Now())
	v.UserNotesCount = 1
	o.Alfa().SetApprovals(7, 3, "echo")
	o.Subscribe("kilo")
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
}

func TestWatchCommentDigest(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Alfa().SeedNote(7, 3, "delta", "looks\ngood", false, time.Now())
	v.UserNotesCount++
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.True(t, strings.HasPrefix(bodies[0], "[group/app!3] "))
	assert.True(t, strings.HasSuffix(bodies[0], " · delta: looks good"))
}

func TestWatchCommentAnnouncedOnce(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Alfa().SeedNote(7, 3, "delta", "first", false, time.Now())
	v.UserNotesCount++
	o.Poll()
	o.Alfa().SeedNote(7, 3, "delta", "second", false, time.Now())
	v.UserNotesCount++
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 2, bodies)
	assert.True(t, strings.HasSuffix(bodies[1], " · delta: second"))
}

func TestWatchBaselineSkipsEarlierNotes(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Alfa().SeedNote(7, 3, "delta", "earlier", false, time.Now())
	v.UserNotesCount = 1
	o.Subscribe("kilo")
	o.Alfa().SeedNote(7, 3, "delta", "later", false, time.Now())
	v.UserNotesCount++
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.True(t, strings.HasSuffix(bodies[0], " · delta: later"))
}

func TestWatchSystemNoteSkipped(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Alfa().SeedNote(7, 3, "delta", "added 1 commit", true, time.Now())
	v.UserNotesCount++
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
}

func TestWatchNotesReadOnlyWhenMoved(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Alfa().SeedNote(7, 3, "delta", "quiet", false, time.Now())
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
	later := v.UpdatedAt.Add(time.Second)
	v.UpdatedAt = &later
	o.Poll()
	assert.Count(t, 1, o.Bodies("kilo"))
	o.Poll()
	assert.Count(t, 1, o.Bodies("kilo"))
}

func TestWatchOwnNoteLeftOut(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Subscribe("lima")
	_, e := o.Service().CreateMergeRequestNote(
		t.Context(),
		constant.FixtureInstance,
		7,
		3,
		"on it",
		"kilo",
	)
	assert.FatalOnError(t, e)
	v.UserNotesCount++
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
	lima := o.Bodies("lima")
	assert.Count(t, 1, lima)
	assert.True(t, strings.HasSuffix(lima[0], " · kilo: on it"))
}

func TestWatchNoteWithoutCallsignShownToAll(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	_, e := o.Service().CreateMergeRequestNote(
		t.Context(),
		constant.FixtureInstance,
		7,
		3,
		"automation",
		"",
	)
	assert.FatalOnError(t, e)
	v.UserNotesCount++
	o.Poll()
	kilo := o.Bodies("kilo")
	assert.Count(t, 1, kilo)
	assert.True(
		t,
		strings.HasSuffix(
			kilo[0],
			join.Empty(" · ", gitlabConstant.MockUser, ": automation"),
		),
	)
}

func TestWatchPipelineAndApproval(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.HeadPipeline.Status = "failed"
	o.Alfa().SetApprovals(7, 3, "echo")
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	lines := strings.Split(bodies[0], "\n")
	assert.True(t, strings.HasPrefix(lines[0], "[group/app!3] 2 new · "))
	assert.Strings(
		t,
		[]string{"    pipeline #10 running → failed", "    echo approved"},
		lines[1:],
	)
}

func TestWatchPushAndNewPipeline(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.SHA = "cccc3333dddd4444"
	v.HeadPipeline = &gitlab.Pipeline{ID: 11, Status: "created"}
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.Strings(
		t,
		[]string{"    pushed cccc3333", "    pipeline #11 created"},
		strings.Split(bodies[0], "\n")[1:],
	)
}

func TestWatchDraftAndMergeability(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.Draft = true
	v.DetailedMergeStatus = "draft_status"
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.Strings(
		t,
		[]string{
			"    marked as draft",
			"    mergeability mergeable → draft_status",
		},
		strings.Split(bodies[0], "\n")[1:],
	)
	v.Draft = false
	v.DetailedMergeStatus = "mergeable"
	o.Poll()
	bodies = o.Bodies("kilo")
	assert.Count(t, 2, bodies)
	assert.StringContains(t, "    marked as ready", bodies[1])
}

func TestWatchTransientMergeStatusSilent(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.DetailedMergeStatus = "checking"
	o.Poll()
	v.DetailedMergeStatus = "mergeable"
	o.Poll()
	v.DetailedMergeStatus = "checking"
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
	v.DetailedMergeStatus = "conflict"
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.True(
		t,
		strings.HasSuffix(bodies[0], " · mergeability mergeable → conflict"),
	)
}

func TestWatchUnapproved(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	o.SeedWatched()
	o.Alfa().SetApprovals(7, 3, "echo")
	o.Subscribe("kilo")
	o.Alfa().SetApprovals(7, 3)
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.True(t, strings.HasSuffix(bodies[0], " · echo unapproved"))
}

func TestWatchMergedEndsWatch(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.State = gitlabConstant.MergedState
	v.MergeUser = &gitlab.BasicUser{Username: "delta"}
	o.Poll()
	notes := o.Connector().Notes()
	assert.Count(t, 2, notes)
	assert.True(t, strings.HasSuffix(notes[0].Body, " · merged by delta"))
	assert.String(t, "immediate", string(notes[0].Delivery))
	assert.String(
		t,
		"[group/app!3] merged by delta, stopped watching",
		notes[1].Body,
	)
	assert.String(t, "queued", string(notes[1].Delivery))
	left, e := o.Service().MergeRequestSubscriptions("kilo")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, left)
	o.Poll()
	assert.Count(t, 2, o.Connector().Notes())
}

func TestWatchClosedWithoutActor(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.State = gitlabConstant.ClosedState
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 2, bodies)
	assert.String(t, "[group/app!3] closed, stopped watching", bodies[1])
}

func TestWatchGoneEndsWatch(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	v.IID = 99
	o.Poll()
	assert.Strings(
		t,
		[]string{"[group/app!3] no longer found, stopped watching"},
		o.Bodies("kilo"),
	)
	left, e := o.Service().MergeRequestSubscriptions("kilo")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, left)
}

func TestWatchRefusedDigestCarriedOver(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Connector().FailFor("kilo", not_found.Format("no session holds kilo"))
	v.HeadPipeline.Status = "failed"
	assert.Count(t, 1, o.Service().PollWatches(t.Context()))
	assert.Count(t, 0, o.Bodies("kilo"))
	o.Connector().FailFor("kilo", nil)
	o.Alfa().SetApprovals(7, 3, "echo")
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.Strings(
		t,
		[]string{"    pipeline #10 running → failed", "    echo approved"},
		strings.Split(bodies[0], "\n")[1:],
	)
}

func TestWatchSurvivesRestart(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	o.Restart()
	v.HeadPipeline.Status = "success"
	o.Poll()
	bodies := o.Bodies("kilo")
	assert.Count(t, 1, bodies)
	assert.True(
		t,
		strings.HasSuffix(bodies[0], " · pipeline #10 running → success"),
	)
}

func TestWatchSubscribeTwice(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	o.SeedWatched()
	o.Subscribe("kilo")
	_, _, existing, e := o.Service().SubscribeMergeRequest(
		t.Context(),
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
		"",
		goclaudedConstant.DeliveryImmediate,
	)
	assert.FatalOnError(t, e)
	assert.True(t, existing)
	all, f := o.Service().MergeRequestSubscriptions("kilo")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, all)
}

func TestWatchSubscribeMergedRefused(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	o.SeedWatched().State = gitlabConstant.MergedState
	_, _, _, e := o.Service().SubscribeMergeRequest(
		t.Context(),
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
		"",
		goclaudedConstant.DeliveryImmediate,
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(
		t,
		"group/app!3 is merged, nothing to watch",
		e.Error(),
	)
}

func TestWatchUnsubscribe(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	removed, e := o.Service().UnsubscribeMergeRequest(
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
	)
	assert.FatalOnError(t, e)
	assert.True(t, removed)
	again, f := o.Service().UnsubscribeMergeRequest(
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
	)
	assert.FatalOnError(t, f)
	assert.False(t, again)
	v.HeadPipeline.Status = "failed"
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
}

func TestWatchResubscribeStartsFresh(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	v := o.SeedWatched()
	o.Subscribe("kilo")
	_, e := o.Service().UnsubscribeMergeRequest(
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		"kilo",
	)
	assert.FatalOnError(t, e)
	v.HeadPipeline.Status = "failed"
	o.Subscribe("kilo")
	o.Poll()
	assert.Count(t, 0, o.Bodies("kilo"))
}

func TestWatchSweepDropsAbsentCallsigns(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	o.SeedWatched()
	o.Subscribe("kilo")
	o.Subscribe("lima")
	o.Connector().Roster = []client.RosterEntry{{Callsign: "lima"}}
	assert.Count(t, 0, o.Service().SweepRoster(t.Context()))
	kilo, e := o.Service().MergeRequestSubscriptions("kilo")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, kilo)
	lima, f := o.Service().MergeRequestSubscriptions("lima")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, lima)
}

func TestWatchSweepKeepsWatchesOnRosterOutage(t *testing.T) {
	o := service_tester.New(t, service_tester.WithWatches())
	o.SeedWatched()
	o.Subscribe("kilo")
	o.Connector().Failure = outage_tester.RefusedDial()
	assert.Count(t, 0, o.Service().SweepRoster(t.Context()))
	kilo, e := o.Service().MergeRequestSubscriptions("kilo")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, kilo)
}
