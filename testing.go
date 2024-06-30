package josh

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
)

var fixtLock = sync.Mutex{}
var fixtCache = make(map[string]map[string]any)

// Call the given function and cache it for the given test.
//
// The best us for it is for "test fixtures" like in pytest.
// Use Fixture to call a helper function and ensure that repeated calls to it
// (from other fixtures) will use the cached value.
//
// The cache is test-local, including sub-tests.
func Fixture[V any](t *testing.T, f func(t *testing.T) V) V {
	t.Helper()
	tname := t.Name()
	fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fixtLock.Lock()
	defer fixtLock.Unlock()
	v, found := fixtCache[tname][fname]
	if !found {
		v = f(t)
		if fixtCache[tname] == nil {
			fixtCache[tname] = make(map[string]any)
		}
		fixtCache[tname][fname] = v
		t.Cleanup(func() {
			fixtLock.Lock()
			defer fixtLock.Unlock()
			delete(fixtCache[tname], fname)
		})
	}
	return v.(V)
}
