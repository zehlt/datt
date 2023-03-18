package datt_test

import (
	"reflect"
	"testing"
)

func AssertEqual[T any](t testing.TB, got T, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNotEqual[T any](t testing.TB, got T, want T) {
	t.Helper()

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertPanic(t testing.TB, fn func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	fn()
}
