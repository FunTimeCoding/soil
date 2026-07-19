package example

type Signal struct {
	done chan struct{}
}

func EmptyStruct() {
	set := map[string]struct{}{}
	set["a"] = struct{}{}
}
