package bitly

func (c *Client) ShortenUrl(longUrl string, domain string, returnType string) (response string, err error) {
	endpointURL := c.BaseURL + "v3/shorten"
	parameters := map[string]string{
		"longUrl": longUrl,
		"domain":  domain,
		"format":  returnType,
	}

	res, err := c.MakeRequest("GET", endpointURL, nil, parameters)
	if err != nil {
		return "", err
	}

	println(string(res))
	return string(res), nil
}
