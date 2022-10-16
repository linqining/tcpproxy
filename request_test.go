package main

import (
	"net"
	"testing"
)

func Test_listenerForBind(t *testing.T) {
	ln, err := tcpListenerForBind()
	t.Log(err)
	spec := ln.Addr().(*net.TCPAddr)
	t.Log(spec.IP)
	t.Log(spec.Port)
}
