package main

import (
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, output interface{}, errMessage ...string) {
	if output != expected {
		t.Error(errMessage)
	}
}