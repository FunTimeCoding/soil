package nut

import (
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"net"
	"strconv"
)

func New(
	host string,
	port int,
) *Client {
	return &Client{
		address: net.JoinHostPort(host, strconv.Itoa(port)),
		timeout: constant.Timeout,
	}
}
