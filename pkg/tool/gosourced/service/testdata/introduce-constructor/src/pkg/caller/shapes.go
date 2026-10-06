package caller

import (
	"example/pkg/single"
	"time"
)

func exact() *single.Shape {
	a := &single.Shape{Draw: 1, Inhale: 2}

	return a
}

func reversed() *single.Shape {
	return &single.Shape{Inhale: 2, Draw: 1}
}

func superset() *single.Shape {
	c := &single.Shape{Draw: 1, Inhale: 2, Hold: 3}

	return c
}

func nestedSuperset() []*single.Shape {
	return []*single.Shape{{Draw: 1, Inhale: 2, Wait: time.Second}}
}

func missing() *single.Shape {
	return &single.Shape{Draw: 1}
}

func value() single.Shape {
	return single.Shape{Draw: 1, Inhale: 2}
}

func ordered() *single.Shape {
	return &single.Shape{Inhale: two(), Draw: one()}
}

func allocated() *single.Shape {
	return new(single.Shape)
}

func one() float64 {
	return 1
}

func two() float64 {
	return 2
}
