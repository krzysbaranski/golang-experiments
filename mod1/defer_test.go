package main

import (
	"testing"
)
func c() (i int) {
	defer func() { i++ }()
	return 1
}


func TestDefer(t *testing.T) {
	if c() != 2 {
		t.Fatal()
	}
}