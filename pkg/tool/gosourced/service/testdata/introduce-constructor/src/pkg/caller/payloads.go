package caller

import "example/pkg/bag"

func payload() *bag.Payload {
	p := &bag.Payload{}

	return p
}

func headers() []*bag.Header {
	return []*bag.Header{
		&bag.Header{Key: "alfa", Value: "bravo"},
		&bag.Header{Key: "charlie"},
	}
}
