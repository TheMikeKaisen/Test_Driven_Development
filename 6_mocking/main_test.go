package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("test countdown", func(t *testing.T) {

		buf := &bytes.Buffer{}

		Countdown(buf)

		got := buf.String()
		want := "321Go!"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
