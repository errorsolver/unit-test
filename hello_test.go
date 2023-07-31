package main

import "testing"

var (
	hello = "Hello"
)

func TestHello(t *testing.T) {
	if hello != "Hello" {
		panic("error test hello")
	}
}
