package main

type Peer interface{}

type Transport interface {
	ListenAndAccept() error
}
