package integers

import "testing"


func TestAdd(t *testing.T){
	t.Run("add two positive numbers:- ",   func(t *testing.T){
		got := add(9, 10)
		want := 19

		tester(t, got, want);
	})

	t.Run("add two negative numbers:- ",   func(t *testing.T){
		got := add(-9, -10)
		want := -19

		tester(t, got, want);
	})
}

func tester(t testing.TB, got, want int){
	t.Helper()

	if got !=  want {
		t.Errorf("got %d, want %d", got , want)
	}

}