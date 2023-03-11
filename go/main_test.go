package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	result := Greeting()
	if result != "Hello World" {
		t.Error("Expected 'Hello World', got ", result)
	}
}
