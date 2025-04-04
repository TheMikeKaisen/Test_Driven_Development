package dictionary

import "testing"

func TestDictionary(t *testing.T) {

	
	t.Run("test using simple Search ftest", func(t *testing.T) {
		
		dict := map[string]string{"test": "this is a test"}
		got, err := Search(dict, "test")
		if err != nil {
			// Fatal/Fatalf will stop the execution if it runs UNLIKE Error / Errorf
			t.Fatalf("error : %s", err.Error())
		}
		want := "this is a test"

		assertError(t, got, want)
	})

	t.Run("test using type method -> when the word is present", func(t *testing.T) {
		dict := Dictionary{"test":"this is a test"}

		// Search is a type method defined in map.go
		got,_ := dict.Search("test")
		want := "this is a test"

		assertError(t, got, want)
	})


	t.Run("test using type method -> when the word is NOT present", func(t *testing.T) {
		dict := Dictionary{"test":"this is a test"}

		// Search is a type method defined in map.go
		_, err := dict.Search("Unknown")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err.Error(), want)
	})


	// -------------------------------ADD ELEMENT TO MAP -----------------------------------------
	t.Run("Adding a value using type method", func(t *testing.T) {
		dict := Dictionary{"test":"value1"}

		dict.Add("test2", "value2")

		got, err := dict.Search("test2")
		if err != nil {
			t.Fatal("key does not exist")
		}

		want := "value2"

		assertError(t, got, want)
	})

}



func assertError(t *testing.T, got, want string){
	t.Helper()

	if got!=want {
		t.Errorf("got %s, want %s", got, want)
	}
}