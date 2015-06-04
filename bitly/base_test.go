package bitly

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShorten(t *testing.T) {
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

func TestExpand(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/expand", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"url": "http://foo.bar"}`)
	})

	res, err := client.Expand([]string{"http://bit.ly/ze6"}, nil, "json")
	want := "{\"url\": \"http://foo.bar\"}"

	if res != want {
		t.Errorf("Returned %v, wanted %v", res, want)
	}

	if err != nil {
		t.Errorf("Error:%v", err)
	}
}

func TestInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/info", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"shortURL": "http://bit.ly/ze6poY"}`)
	})

	res, err := client.Info([]string{"http://foo.bar"}, nil, true)
	want := "{\"shortURL\": \"http://bit.ly/ze6poY\"}"

	if res != want {
		t.Errorf("Returned %v, wanted %v", res, want)
	}

	if err != nil {
		t.Errorf("Error:%v", err)
	}
}
