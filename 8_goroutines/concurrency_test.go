package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsites(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestConcurrency(t *testing.T) {
	t.Run("testing concurrency", func(t *testing.T) {

		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		got := WebsiteChecker(mockWebsites, websites)

		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want);
		}

	})
}
