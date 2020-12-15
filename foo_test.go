package main

import "testing"

func TestFoo(t *testing.T) {
	want := "Bar"
	if got := Foo(); got != want {
		t.Errorf("Foo() = %v, want %v", got, want)
	}
}
