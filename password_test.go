package main

import "testing"

func TestHashPassword(t *testing.T) {
	p, _ := HashPassword("123456")

	if !ComparePassword(p, "123456") {
		t.Fatal()
	}
}
