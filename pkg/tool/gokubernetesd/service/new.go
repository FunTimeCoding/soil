package service

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"

func New(s *store.Store) *Service {
	return &Service{
		clusters: buildClusters(),
		store:    s,
	}
}
