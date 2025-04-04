package di

import (
	"bytes"
	"testing"
)



func TestLogger(t *testing.T) {
	t.Run("testing io.writer by injecting bytes.buffer", func(t *testing.T) {
		var buf bytes.Buffer

		// injecting bytes.Buffer as a dependency
		err := Logger(&buf, "karthik")
		if err != nil {
			t.Errorf("%q", err)
		}

		got := buf.String()
		want := "Hello, karthik"

		if got!= want {
			t.Errorf("got %s, want %s", got, want)
		}


		
	})
}