package joshtest

import "testing"

// Replace replaces the value of target with val.
//
// The old value is restored when the test ends.
func Replace[T any](t testing.TB, target *T, val T) {
	t.Helper()
	if target == nil {
		t.Fatalf("Replace: nil pointer")
		panic("unreachable") // pacify staticcheck
	}
	old := *target
	t.Cleanup(func() {
		*target = old
	})

	*target = val
}
