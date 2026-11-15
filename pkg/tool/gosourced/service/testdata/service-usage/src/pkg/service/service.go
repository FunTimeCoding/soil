package service

import "example/pkg/store"

type Service struct {
	store *store.Store
	clock func() int
}
