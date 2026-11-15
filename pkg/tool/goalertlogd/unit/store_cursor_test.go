package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/store_tester"
	"testing"
)

func TestAfterWalksForwardInIdentifierOrder(t *testing.T) {
	s := store_tester.New(t)
	store_tester.Create(t, s, 3)
	first, e := s.After(0, 2)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, first)
	assert.Integer(t, 1, int(first[0].Identifier))
	assert.Integer(t, 2, int(first[1].Identifier))
	rest, f := s.After(first[1].Identifier, 2)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, rest)
	assert.Integer(t, 3, int(rest[0].Identifier))
	done, g := s.After(rest[0].Identifier, 2)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, done)
}

func TestByIdentifiersSkipsUnknownOnes(t *testing.T) {
	s := store_tester.New(t)
	store_tester.Create(t, s, 3)
	result, e := s.ByIdentifiers([]uint{3, 1, 99})
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
	assert.Integer(t, 1, int(result[0].Identifier))
	assert.Integer(t, 3, int(result[1].Identifier))
}

func TestLatestIdentifierIsHead(t *testing.T) {
	s := store_tester.New(t)
	empty, e := s.LatestIdentifier()
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, int(empty))
	store_tester.Create(t, s, 3)
	latest, f := s.LatestIdentifier()
	assert.FatalOnError(t, f)
	assert.Integer(t, 3, int(latest))
}
