package bitly

import (
	"fmt"
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
	client.BaseURL = server.URL
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		print(r.URL.String())
	})

	parameters := map[string]string{
		"foo": "bar",
	}

	response, err := client.MakeRequest("GET", "/", "asd", parameters)
	if err != nil {
		t.Errorf("Error:%v", err)
	}
	fmt.Println(string(response))
}
