package main

import (
	"io"
	"net/http"
	"net/url"
)

type HTTPHeaders map[string]string

type httpClient struct {
	writeURL    string
	HTTPHeaders HTTPHeaders
	UserAgent   string
	Username    string
	Password    string
	// client   *http.Client
	url *url.URL
}

func (c *httpClient) makeRequest(uri string, body io.Reader) (*http.Request, error) {
	var req *http.Request
	var err error
	if c.config.ContentEncoding == "gzip" {
		body, err = compressWithGzip(body)
		if err != nil {
			return nil, err
		}
	}
	req, err = http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	for header, value := range c.HTTPHeaders {
		req.Header.Set(header, value)
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	return req, nil
}

func main() {
	test := "http://123.0.0.0:1"
	url1, _ := url.Parse(test)
	player1 := httpClient{
		writeURL:    "neo4j",
		HTTPHeaders: nil,
		UserAgent:   "neo4j",
		Password:    "na",
		url:         url1,
	}
	makeRequest(player1.url, 123)
}
