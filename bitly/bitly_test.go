package bitly

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient("abc", "")
	client.BaseURL = server.URL + "/"
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	c := NewClient("abc", "")

	if got, want := c.AccessToken, "abc"; got != want {
		t.Errorf("NewClient AccessToken is %v, want %v", got, want)
	}
}

func TestMakeRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		if p := "bar"; p != r.URL.Query().Get("foo") {
			t.Errorf("Parameter = %v, want %v", r.URL.Query().Get("foo"), p)
		}
	})

	_, err := client.MakeRequest("GET", client.BaseURL+"/foo?foo=bar", "asd")
	if err != nil {
		t.Errorf("Error:%v", err)
	}
}
