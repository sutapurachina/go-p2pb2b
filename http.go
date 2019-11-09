package p2pb2b

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type auth struct {
	ServiceToken string
	AccessToken  string
}

type client struct {
	http *http.Client
	auth *auth
}

type response struct {
	Header     http.Header
	Body       io.ReadCloser
	StatusCode int
	Status     string
}

func mergeHeaders(firstHeaders map[string]string, secondHeaders map[string]string) map[string]string {
	if secondHeaders == nil {
		return firstHeaders
	}
	if firstHeaders == nil {
		return secondHeaders
	}
	for k, v := range secondHeaders {
		if firstHeaders[k] == "" {
			firstHeaders[k] = v
		}
	}
	return firstHeaders
}

func query(params url.Values) string {
	cleanedParams := url.Values{}
	for k, v := range params {
		if params[k][0] != "" {
			cleanedParams[k] = v
		}
	}
	if cleanedParams == nil || len(cleanedParams) == 0 {
		return ""
	}
	return fmt.Sprintf("?%s", cleanedParams.Encode())
}

func (c *client) sendPost(url string, additionalHeaders map[string]string, body io.Reader) (*response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return &response{}, fmt.Errorf("error creating POST request, %v", err)
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendPut(url string, additionalHeaders map[string]string, body io.Reader) (*response, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return &response{}, fmt.Errorf("error creating PUT request, %v", err)
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendGet(url string, additionalHeaders map[string]string) (*response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &response{}, fmt.Errorf("error creating GET request, %v", err)
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendHead(url string, additionalHeaders map[string]string) (*response, error) {
	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return &response{}, fmt.Errorf("error creating HEAD request, %v", err)
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendDelete(url string, additionalHeaders map[string]string) (*response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return &response{}, fmt.Errorf("error creating DELETE request, %v", err)
	}

	return c.sendRequest(req, additionalHeaders)
}

func (c *client) sendRequest(request *http.Request, additionalHeaders map[string]string) (*response, error) {

	for k, v := range additionalHeaders {
		request.Header.Add(k, v)
	}

	thisHeaders := map[string]string{}
	if c.auth != nil {
		thisHeaders["Authorization"] = fmt.Sprintf("bearer %s", c.auth.ServiceToken)
		thisHeaders["X-User-Token"] = fmt.Sprintf("bearer %s", c.auth.AccessToken)
	}

	headers := mergeHeaders(additionalHeaders, thisHeaders)

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	resp, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}

	return &response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Header:     resp.Header,
		Body:       resp.Body,
	}, nil
}
