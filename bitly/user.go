package bitly

import (
	"net/url"
)

func (c *Client) LinkLookup(urls []string) (response string, err error) {
	endpointURL, err := url.Parse(c.BaseURL + "v3/user/link_lookup")

	if err != nil {
		return "", err
	}
	params := url.Values{}
	for _, url := range urls {
		params.Add("url", url)
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

//func (c *Client) LinkEdit(url string, title string, note string, private bool) (response string, err error) {
//endpointURL, err := url.Parse(c.BaseURL + "v3/user/link_edit")
//}
