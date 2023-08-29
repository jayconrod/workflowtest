package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	got := os.Getenv("EXAMPLE_SECRET")
	if got == "" {
		t.Fatalf("EXAMPLE_SECRET not set")
	}
	t.Logf("the secret is: %s", got)
	want := "EXAMPLE_SECRET"
	if got != want {
		t.Fatalf("EXAMPLE_SECRET has unexpected value %q", got)
	}
}
