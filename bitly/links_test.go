package bitly

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShortenURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/shorten", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"shortURL": "http://bit.ly/ze6poY"}`)
	})

	res, err := client.Shorten("http://foo.bar", "bit.ly", "json")
	want := "{\"shortURL\": \"http://bit.ly/ze6poY\"}"

	if res != want {
		t.Errorf("Returned %v, wanted %v", res, want)
	}

	if err != nil {
		t.Errorf("Error:%v", err)
	}
}
