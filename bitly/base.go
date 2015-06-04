package bitly

import (
	"net/url"
)

func (c *Client) Shorten(longURL string, shortDomain string, returnFormat string) (response string, err error) {
	endpointURL, err := url.Parse(c.BaseURL + "v3/shorten")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("longUrl", longURL)
	if shortDomain != "" {
		params.Set("domain", shortDomain)
	}
	if returnFormat == "" {
		params.Set("format", "json")
	} else {
		params.Set("format", returnFormat)
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) Expand(shortURLs []string, hashes []string, returnFormat string) (response string, err error) {
	endpointURL, err := url.Parse(c.BaseURL + "v3/expand")

	if err != nil {
		return "", err
	}
	params := url.Values{}
	for _, url := range shortURLs {
		params.Add("shortUrl", url)
	}
	for _, hash := range hashes {
		params.Add("hash", hash)
	}

	if returnFormat == "" {
		params.Set("format", "json")
	} else {
		params.Set("format", returnFormat)
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) Info(shortURLs []string, hashes []string, extraUserInfo bool) (response string, err error) {
	endpointURL, err := url.Parse(c.BaseURL + "v3/info")

	if err != nil {
		return "", err
	}
	params := url.Values{}
	for _, url := range shortURLs {
		params.Add("shortUrl", url)
	}
	for _, hash := range hashes {
		params.Add("hash", hash)
	}
	if extraUserInfo {
		params.Add("expand_user", "true")
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
