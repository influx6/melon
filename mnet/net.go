// Package mnet implements function composition for net.Listeners and net.Conns providing functional compositions
// where behaviours are hidden within functions and reducing alot of deadlocks.
package mnet

import (
	"crypto/tls"
	"errors"
	"net"
	"sync"

	"github.com/influx6/faux/netutils"
	"github.com/influx6/melon"
)

// errors ...
var (
	ErrListenerClosed = errors.New("listener has being closed")
)

// Listen returns a Readr and Writer for reading net.Conns from a underline net.Listener.
func Listen(protocol string, addr string, config *tls.Config) (melon.ConnReadWriteCloser, error) {
	lt, err := netutils.MakeListener(protocol, addr, config)
	if err != nil {
		return nil, err
	}

	if tlt, ok := lt.(*net.TCPListener); ok {
		lt = netutils.NewKeepAliveListener(tlt)
	}

	readWriter := new(connReadWriter)
	readWriter.l = lt

	return readWriter, nil
}

type connReadWriter struct {
	ml sync.Mutex
	l  net.Listener
}

// WriteConn receives the provided net.Conn and closes the connection, this
// assumes all operation with the net.Conn has being complete and the resource
// and connection should end here.
func (cs *connReadWriter) WriteConn(conn net.Conn) error {
	return conn.Close()
}

// Close closes the underneath net.Listener, ending all
func (cs *connReadWriter) Close() error {
	cs.ml.Lock()
	defer cs.ml.Unlock()

	if cs.l == nil {
		return ErrListenerClosed
	}

	err := cs.l.Close()
	cs.l = nil
	return err
}

// ReadConn returns a new net.Conn from the underying listener.
func (cs *connReadWriter) ReadConn() (net.Conn, error) {
	cs.ml.Lock()
	defer cs.ml.Unlock()

	if cs.l == nil {
		return nil, ErrListenerClosed
	}

	newConn, err := cs.l.Accept()
	if err != nil {
		return nil, err
	}

	return newConn, nil
}
