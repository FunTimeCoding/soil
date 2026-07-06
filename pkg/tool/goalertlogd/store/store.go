package store

import "github.com/funtimecoding/soil/pkg/bolt"

type Store struct {
	client *bolt.Client
}
