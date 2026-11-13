package status

type Status struct {
	Collections []CollectionStatus
}

type CollectionStatus struct {
	Name string
}
