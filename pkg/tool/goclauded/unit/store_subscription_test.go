package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

func TestSubscribeStartsAtZero(t *testing.T) {
	s := store_tester.New(t)
	result, e := s.Store.Subscribe("consumer", []string{constant.Label})
	assert.FatalOnError(t, e)
	assert.String(t, "consumer", result.Name)
	assert.Integer(t, 0, int(result.EventIdentifier))
	assert.Count(t, 1, result.KindList())
	assert.String(t, "label", result.KindList()[0])
}

func TestSubscribeIsIdempotent(t *testing.T) {
	s := store_tester.New(t)
	_, e := s.Store.Subscribe("consumer", []string{constant.Label})
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.Store.SetSubscriptionPosition("consumer", 42))
	again, e := s.Store.Subscribe("consumer", []string{constant.Announce})
	assert.FatalOnError(t, e)
	assert.Integer(t, 42, int(again.EventIdentifier))
	assert.String(t, "announce", again.KindList()[0])
}

func TestSubscriptionByNameUnknown(t *testing.T) {
	s := store_tester.New(t)
	_, e := s.Store.SubscriptionByName("nobody")
	assert.NotNil(t, e)
}

func TestSetSubscriptionPositionAdvances(t *testing.T) {
	s := store_tester.New(t)
	_, e := s.Store.Subscribe("consumer", nil)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.Store.SetSubscriptionPosition("consumer", 7))
	result, e := s.Store.SubscriptionByName("consumer")
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, int(result.EventIdentifier))
}

func TestSetSubscriptionPositionAllowsReplay(t *testing.T) {
	s := store_tester.New(t)
	_, e := s.Store.Subscribe("consumer", nil)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.Store.SetSubscriptionPosition("consumer", 7))
	assert.FatalOnError(t, s.Store.SetSubscriptionPosition("consumer", 3))
	result, e := s.Store.SubscriptionByName("consumer")
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, int(result.EventIdentifier))
}

func TestSetSubscriptionPositionUnknown(t *testing.T) {
	s := store_tester.New(t)
	assert.NotNil(t, s.Store.SetSubscriptionPosition("nobody", 1))
}

func TestSubscribeWithoutKinds(t *testing.T) {
	s := store_tester.New(t)
	result, e := s.Store.Subscribe("consumer", nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result.KindList())
}
