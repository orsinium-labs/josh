package joshtest_test

import (
	"math/rand"
	"testing"

	"github.com/orsinium-labs/josh/joshtest"
)

func getRand(*testing.T) int {
	return rand.Int()
}

func getRand2(*testing.T) int {
	return rand.Int()
}

func TestFixture(t *testing.T) {
	r1 := joshtest.Fixture(t, getRand)
	r2 := joshtest.Fixture(t, getRand)
	r3 := joshtest.Fixture(t, getRand2)
	if r1 != r2 {
		t.FailNow()
	}
	if r1 == r3 {
		t.FailNow()
	}
}
