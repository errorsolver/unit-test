package main

import "testing"

var (
	mockupName = "Bulan"
	salam      = "Salam Enigma"
)

func TestGreeting(t *testing.T) {
	if mockupName != "Bulan" {
		panic("Wrong Name, should be bulan")
	}
}

func TestSaySalam(t *testing.T) {
	if salam != "Salam Enigma" {
		panic("wrong salam")
	}
}
