package josh

import (
	"reflect"
	"sync"
	"testing"
)

var fixtLock = sync.Mutex{}
var fixtCache = make(map[string]map[uintptr]any)

// Call the given function and cache it for the given test.
//
// The best us for it is for "test fixtures" like in pytest.
// Use Fixture to call a helper function and ensure that repeated calls to it
// (from other fixtures) will use the cached value.
//
// The cache is test-local, including sub-tests.
//
//	user := josh.Fixture(t, GetUser)
func Fixture[V any](t *testing.T, f func(t *testing.T) V) V {
	t.Helper()
	tname := t.Name()
	fptr := reflect.ValueOf(f).Pointer()
	fixtLock.Lock()
	defer fixtLock.Unlock()
	v, found := fixtCache[tname][fptr]
	if !found {
		v = f(t)
		if fixtCache[tname] == nil {
			fixtCache[tname] = make(map[uintptr]any)
		}
		fixtCache[tname][fptr] = v
		t.Cleanup(func() {
			fixtLock.Lock()
			defer fixtLock.Unlock()
			delete(fixtCache[tname], fptr)
		})
	}
	return v.(V)
}
