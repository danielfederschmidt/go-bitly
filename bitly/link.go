package bitly

import (
	"net/url"
	"time"
)

func (c *Client) LinkClicks(link string, timeUnit string, timeValue int, utcOffset int,
	rollup bool, limit int, unitReferenceTimestamp *time.Time, returnFormat string) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/clicks")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)
	params.Set("unit", timeUnit)
	params.Set("units", string(timeValue))
	params.Set("tz_offset", string(utcOffset))

	if rollup {
		params.Set("rollup", "true")
	} else {
		params.Set("rollup", "false")
	}
	params.Set("limit", string(limit))

	params.Set("format", returnFormat)

	if unitReferenceTimestamp != nil {
		params.Set("unit_reference_ts", unitReferenceTimestamp.String())
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) LinkCountries(link string, timeUnit string, timeValue int, utcOffset int, limit int, unitReferenceTimestamp *time.Time) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/countries")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)
	params.Set("unit", timeUnit)
	params.Set("units", string(timeValue))
	params.Set("tz_offset", string(utcOffset))

	params.Set("limit", string(limit))

	if unitReferenceTimestamp != nil {
		params.Set("unit_reference_ts", unitReferenceTimestamp.String())
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) LinkEncoders(link string, restrictNetwork bool, limit int, expandUser bool) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/encoders")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)

	if restrictNetwork {
		params.Set("my_network", "true")
	} else {
		params.Set("my_network", "false")
	}

	params.Set("limit", string(limit))

	if expandUser {
		params.Set("expand_user", "true")
	} else {
		params.Set("expand_user", "false")
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) LinkEncodersByCount(link string, restrictNetwork bool, limit int, expandUser bool) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/encoders_by_count")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)

	if restrictNetwork {
		params.Set("my_network", "true")
	} else {
		params.Set("my_network", "false")
	}
	params.Set("limit", string(limit))

	if expandUser {
		params.Set("expand_user", "true")
	} else {
		params.Set("expand_user", "false")
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) LinkEncodeCount(link string) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/encoders_count")

	params := url.Values{}
	params.Set("link", link)

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

func (c *Client) LinkLookup(links []string) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/lookup")

	params := url.Values{}
	for _, link := range links {
		params.Add("url", link)
	}

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil

}
func (c *Client) LinkInfo(link string) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/info")

	params := url.Values{}
	params.Set("link", link)

	reqURL := endpointURL.String() + "?" + params.Encode()
	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

func (c *Client) LinkReferrers(link string, timeUnit string, timeValue int, utcOffset int, limit int, unitReferenceTimestamp *time.Time) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/referrers")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)
	params.Set("unit", timeUnit)
	params.Set("units", string(timeValue))
	params.Set("tz_offset", string(utcOffset))

	params.Set("limit", string(limit))

	if unitReferenceTimestamp != nil {
		params.Set("unit_reference_ts", unitReferenceTimestamp.String())
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) LinkReferrersByDomain(link string, timeUnit string, timeValue int, utcOffset int, limit int, unitReferenceTimestamp *time.Time) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/referrers_by_domain")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)
	params.Set("unit", timeUnit)
	params.Set("units", string(timeValue))
	params.Set("tz_offset", string(utcOffset))

	params.Set("limit", string(limit))

	if unitReferenceTimestamp != nil {
		params.Set("unit_reference_ts", unitReferenceTimestamp.String())
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) ReferreringDomains(link string, timeUnit string, timeValue int, utcOffset int, limit int, unitReferenceTimestamp *time.Time) (response string, err error) {

	endpointURL, err := url.Parse(c.BaseURL + "v3/link/referrers_by_domain")

	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("link", link)
	params.Set("unit", timeUnit)
	params.Set("units", string(timeValue))
	params.Set("tz_offset", string(utcOffset))

	params.Set("limit", string(limit))

	if unitReferenceTimestamp != nil {
		params.Set("unit_reference_ts", unitReferenceTimestamp.String())
	}

	reqURL := endpointURL.String() + "?" + params.Encode()

	res, err := c.MakeRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
