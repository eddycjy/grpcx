package driver

import (
	"net"
)

type DriverIface interface {
	Serve() error
	SetListener(net.Listener)
}