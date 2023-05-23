package main

import "testing"

func TestInitialize(t *testing.T) {
	err := initialize()
	if err != nil {
		t.Error("main failed to initialize")
	}
}
