package main

import "testing"

func TestGreetsByName(t *testing.T) {
	_, greeting, err := greet(t.Context(), nil, GreetInput{Name: "Ada"})
	if err != nil {
		t.Fatal(err)
	}
	if greeting.Message != "Hello, Ada!" {
		t.Errorf("greeted with %q", greeting.Message)
	}
}

func TestAnEmptyNameIsAToolError(t *testing.T) {
	if _, _, err := greet(t.Context(), nil, GreetInput{}); err == nil {
		t.Error("greeted no one")
	}
}
