package main

import "testing"

func TestSendMessage(t *testing.T) {
	expected := "Hello, CI/CD!111"
	result := SendMessage()
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
