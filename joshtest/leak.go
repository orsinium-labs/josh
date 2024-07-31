package joshtest

import (
	"runtime"
	"testing"
	"time"
)

// Call at the beginning of a test to ensure that there are no goroutine leak.
func DetectGoLeak(t testing.TB) {
	t.Helper()
	startN := runtime.NumGoroutine()
	t.Cleanup(func() {
		if t.Failed() {
			// Something else went wrong.
			return
		}
		// Goroutines might be still exiting.
		for i := 0; i < 100; i++ {
			if runtime.NumGoroutine() <= startN {
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
		endN := runtime.NumGoroutine()
		if endN <= startN {
			return
		}
		t.Fatalf("goroutine count: expected %d, got %d\n", startN, endN)
	})
}
