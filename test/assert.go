package test

import "testing"

func Assert(t *testing.T, cond bool) {
	if !cond {
		t.Errorf("condition failed")
	}
}
