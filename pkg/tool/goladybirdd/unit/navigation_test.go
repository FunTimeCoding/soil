package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/unit/service_tester"
	"testing"
	"time"
)

func TestClickWaitsForNavigationLoad(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.NavigateOn(
		constant.ClickScript,
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/next",
		),
		o.Event(
			browserConstant.LadybirdLoadEvent,
			"alfa",
			"https://alfa.example/next",
		),
	)
	v, e := o.Service().Click("", "e1", true)
	assert.FatalOnError(t, e)
	assert.String(t, "https://alfa.example/next", v.Landed)
	assert.False(t, v.Loading)
}

func TestClickWithoutNavigationLandsNowhere(t *testing.T) {
	v, e := service_tester.New(t).Clickable().Service().Click("", "e1", false)
	assert.FatalOnError(t, e)
	assert.String(t, "", v.Landed)
	assert.False(t, v.Loading)
}

func TestClickNavigationStillOpenLoading(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.NavigateOn(
		constant.ClickScript,
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/slow",
		),
	)
	v, e := o.Service().Click("", "e1", false)
	assert.FatalOnError(t, e)
	assert.True(t, v.Loading)
}

func TestClickHistoryUpdateLanded(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.NavigateOn(
		constant.ClickScript,
		o.Event(
			browserConstant.LadybirdHistoryUpdatedEvent,
			"alfa",
			"https://alfa.example/routed",
		),
	)
	v, e := o.Service().Click("", "e1", false)
	assert.FatalOnError(t, e)
	assert.String(t, "https://alfa.example/routed", v.Landed)
	assert.False(t, v.Loading)
}

func TestClickFailedNavigationSettled(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.NavigateOn(
		constant.ClickScript,
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/gone",
		),
		o.Event(
			browserConstant.LadybirdNavigationFailedEvent,
			"alfa",
			"https://alfa.example/gone",
		),
	)
	v, e := o.Service().Click("", "e1", false)
	assert.FatalOnError(t, e)
	assert.False(t, v.Loading)
}

func TestSnapshotWaitsOutOpenNavigation(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.Emit(
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/slow",
		),
	)
	v, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	assert.True(t, v.Loading)
}

func TestScreenshotWaitsOutOpenNavigation(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.Emit(
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/slow",
		),
	)
	start := time.Now()
	_, e := o.Service().Screenshot("")
	assert.FatalOnError(t, e)
	assert.GreaterEqual(t, 300*time.Millisecond, time.Since(start))
}

func TestEvaluateWaitsOutOpenNavigation(t *testing.T) {
	o := service_tester.New(t).Clickable()
	o.Client.Emit(
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"alfa",
			"https://alfa.example/slow",
		),
	)
	start := time.Now()
	_, e := o.Service().Evaluate("", "1+2")
	assert.FatalOnError(t, e)
	assert.GreaterEqual(t, 300*time.Millisecond, time.Since(start))
}

func TestWatcherKeepsTabsApart(t *testing.T) {
	o := service_tester.New(t)
	w := navigation.New(50*time.Millisecond, 50*time.Millisecond)
	mark := w.Mark("alfa")
	w.Observe(
		o.Event(
			browserConstant.LadybirdNavigationStartedEvent,
			"bravo",
			"https://bravo.example/",
		),
	)
	assert.False(t, w.StartedSince("alfa", mark))
	assert.True(t, w.StartedSince("bravo", 0))
}

func TestWatcherWakesOnLateStart(t *testing.T) {
	o := service_tester.New(t)
	w := navigation.New(time.Second, time.Second)
	mark := w.Mark("alfa")
	go func() {
		time.Sleep(20 * time.Millisecond)
		w.Observe(
			o.Event(
				browserConstant.LadybirdNavigationStartedEvent,
				"alfa",
				"https://alfa.example/",
			),
		)
	}()
	assert.True(t, w.StartedSince("alfa", mark))
}
