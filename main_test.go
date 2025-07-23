package main

import "testing"

func TestSendMessage(t *testing.T) {
	expected := "Hello, CI/CD!"
	result := SendMessage()
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
