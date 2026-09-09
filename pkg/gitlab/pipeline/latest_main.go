package pipeline

func LatestMain(
	v []*Pipeline,
	mainHash string,
) *Pipeline {
	for _, e := range v {
		if e.Hash == mainHash {
			return e
		}
	}

	return nil
}
