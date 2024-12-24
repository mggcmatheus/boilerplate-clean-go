package main

import "testing"

// TestHelloWorld Testa a função HelloWorld
func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	got := HelloWorld()

	if got != expected {
		t.Errorf("HelloWorld() = %v, want %v", got, expected)
	}
}
