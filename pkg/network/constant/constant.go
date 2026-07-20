package constant

import "github.com/funtimecoding/soil/pkg/network"

const LocalhostAddressString = "127.0.0.1"
const NullPhysicalAddressString = "00:00:00:00:00:00"

var NullPhysicalAddress = network.PhysicalAddress(NullPhysicalAddressString)
