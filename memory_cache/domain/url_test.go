package domain

import "testing"

func TestUrlShorten(t *testing.T) {
	url := Url{Original: "http://test.com:8080/haha"}

	url.Shorten()

	got := url.Shortened
	want := "http://localhost:8080/shortened"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
