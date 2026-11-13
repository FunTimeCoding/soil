package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"testing"
)

func TestFaceFingerprintIgnoresOrder(t *testing.T) {
	alfa := fact.NewInterface(
		"alfa",
		"Bravo",
		map[string]string{"Charlie": "func()"},
	)
	delta := fact.NewInterface(
		"delta",
		"Echo",
		map[string]string{"Foxtrot": "func() int"},
	)
	first := face.New(nil)
	first.Add([]*fact.Interface{alfa, delta})
	second := face.New(nil)
	second.Add([]*fact.Interface{delta, alfa})
	assert.String(t, first.Fingerprint(), second.Fingerprint())
}

func TestFaceFingerprintMovesWithSignature(t *testing.T) {
	first := face.New(nil)
	first.Add(
		[]*fact.Interface{
			fact.NewInterface(
				"alfa",
				"Bravo",
				map[string]string{"Charlie": "func()"},
			),
		},
	)
	second := face.New(nil)
	second.Add(
		[]*fact.Interface{
			fact.NewInterface(
				"alfa",
				"Bravo",
				map[string]string{"Charlie": "func() error"},
			),
		},
	)
	assert.True(t, first.Fingerprint() != second.Fingerprint())
}
