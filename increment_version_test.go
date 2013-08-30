package main

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	check := func(before version, after version) {
		before.increment()
		if before.Full != after.Full {
			t.Fatalf("Expected %s, got %s", after.Full, before.Full)
		}
	}
	check(newVersion(0, 0, 0), newVersion(0, 0, 1))
	check(newVersion(0, 0, 99), newVersion(0, 1, 0))
	check(newVersion(0, 99, 99), newVersion(1, 0, 0))
	check(newVersion(99, 99, 99), newVersion(100, 0, 0))
	check(newVersion(0, 0, 1), newVersion(0, 0, 2))
	check(newVersion(0, 1, 99), newVersion(0, 2, 0))
	check(newVersion(1, 0, 99), newVersion(1, 1, 0))
}
