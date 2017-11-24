package net_test

import (
	"net"
	"testing"
	"time"

	"github.com/influx6/faux/tests"
	mnet "github.com/influx6/melon/net"
)

func TestReadWriter(t *testing.T) {
	reader, err := mnet.Listen("tcp", ":4050", nil)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully created reader and writer")
	}
	tests.Passed("Should have successfully created reader and writer")

	defer reader.Close()

	go makeConn(":4050")
	conn, err := reader.ReadConn()
	if err != nil {
		tests.FailedWithError(err, "Should have sucessfully read connection")
	}
	tests.Passed("Should have sucessfully read connection")

	reader.WriteConn(conn)
}

func makeConn(addr string) {
	conn, err := net.DialTimeout("tcp", addr, 30*time.Second)
	if err != nil {
		return
	}

	conn.Close()
}
