package bitly

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (c *Client) MakeRequest(method, urlStr string, body interface{},
	parameters map[string]string) (res []byte, err error) {

	endpointUrl, err := url.Parse(urlStr)

	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("access_token", c.AccessToken)
	for k, v := range parameters {
		params.Add(k, v)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	reqUrl := endpointUrl.String() + "?" + params.Encode()

	req, err := http.NewRequest(method, reqUrl, buf)

	httpResponse, err := c.Client.Do(req)

	//Todo: Decode Response according to requested Return Type
	jsonResponseData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}
	return jsonResponseData, nil
}
