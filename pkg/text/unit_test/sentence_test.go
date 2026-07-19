package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/sentence"
	"testing"
)

func TestSentence(t *testing.T) {
	one := sentence.New("A")
	assert.String(t, "A.", one.Join())
	two := sentence.New("A")
	two.Add("B")
	assert.String(t, "A, and B.", two.Join())
	three := sentence.New("A")
	three.Add("B")
	three.Add("C")
	assert.String(t, "A, B, and C.", three.Join())
	four := sentence.New("A")
	four.Add("B")
	four.Add("C")
	four.Add("D")
	assert.String(t, "A, B, C, and D.", four.Join())
}
