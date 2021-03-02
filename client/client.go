package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	address string
	baseURL *url.URL
}

func NewClient(host string, port int) (*Client, error) {
	address := fmt.Sprintf("http://%s:%v", host, port)
	baseURL, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("parsing URL `%s`: %w", address, err)
	}

	return &Client{
		address: address,
		baseURL: baseURL,
	}, nil
}

// getBaseURL returns a copy of baseURL. we cannot use copy since there are nested structs in URL
func (c *Client) getBaseURL() (*url.URL, error) {
	baseURL, err := url.Parse(c.baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing URL `%s`: %w", c.baseURL.String(), err)
	}

	return baseURL, nil
}

func (c *Client) Get(key string) ([]byte, error) {
	getURL, err := c.getBaseURL()
	if err != nil {
		return nil, fmt.Errorf("getting base URL: %w", err)
	}

	getURL.Path = "get"
	getURL.RawQuery = fmt.Sprintf("key=%s", url.QueryEscape(key))

	resp, err := http.Get(getURL.String())
	if err != nil {
		return nil, fmt.Errorf("making request to `%s`: %w", getURL.String(), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got non 200 status requesting `%s`: %v message: `%s`", getURL.String(), resp.StatusCode, body)
	}

	return body, nil
}

func (c *Client) Set(key string, value []byte) error {
	setURL, err := c.getBaseURL()
	if err != nil {
		return fmt.Errorf("getting base URL: %w", err)
	}

	setURL.Path = "set"
	setURL.RawQuery = fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(string(value)))

	resp, err := http.Get(setURL.String())
	if err != nil {
		return fmt.Errorf("making request to `%s`: %w", setURL.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("reading response body: %w", err)
		}

		return fmt.Errorf("got non 200 status requesting `%s`: %v message: `%s`", setURL.String(), resp.StatusCode, body)
	}

	return nil
}
