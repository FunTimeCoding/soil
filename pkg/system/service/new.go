package service

import "github.com/funtimecoding/soil/pkg/system"

func New() *Client {
	return &Client{home: system.Home()}
}
