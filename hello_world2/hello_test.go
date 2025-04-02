package helloworld2

import "testing"

func TestHello(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := Hello("chris")
		want := "Hello, chris"

		assertTest(t, got, want)
	})
	t.Run("test if no name is provided!", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertTest(t, got, want)
	})
}

func assertTest(t testing.TB, got, want string) {
	t.Helper() // tells the program that this is a helper function only!
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
