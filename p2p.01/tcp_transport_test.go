package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tcpTransport := NewTCPTransport(listenAddr)

	assert.Equal(t, tcpTransport.listenAddr, listenAddr)

	assert.Nil(t, tcpTransport.ListenAndAccept())
}
