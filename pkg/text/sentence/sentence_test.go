package sentence

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSentence(t *testing.T) {
	one := New("A")
	assert.String(t, "A.", one.Join())
	two := New("A")
	two.Add("B")
	assert.String(t, "A, and B.", two.Join())
	three := New("A")
	three.Add("B")
	three.Add("C")
	assert.String(t, "A, B, and C.", three.Join())
	four := New("A")
	four.Add("B")
	four.Add("C")
	four.Add("D")
	assert.String(t, "A, B, C, and D.", four.Join())
}
