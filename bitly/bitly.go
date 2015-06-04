package bitly

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	defaultBaseURL = "https://api-ssl.bitly.com/"
)

type Client struct {
	Client      *http.Client
	BaseURL     string
	AccessToken string
}

func NewClient(accessToken string, baseURL string) *Client {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	client := &Client{Client: http.DefaultClient, AccessToken: accessToken, BaseURL: baseURL}

	return client
}

// Makes a request to an endpoint and returns the response of the request
// urlStr is the relative URL of the endpoint
// Method is a String contained a HTTP Method
func (c *Client) MakeRequest(method, reqURLString string, body interface{}) (res []byte, err error) {

	//Inject AccessToken
	reqURL := reqURLString + "&access_token=" + c.AccessToken

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, reqURL, buf)

	if err != nil {
		return nil, err
	}

	httpResponse, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}
	return responseData, nil
}
