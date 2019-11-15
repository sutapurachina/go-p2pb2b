package p2pb2b

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const HEADER_X_TXC_APIKEY = "X-TXC-APIKEY"
const HEADER_X_TXC_PAYLOAD = "X-TXC-PAYLOAD"
const HEADER_X_TXC_SIGNATURE = "X-TXC-SIGNATURE"

type auth struct {
	APIKey    string
	APISecret string
}

type client struct {
	http *http.Client
	auth *auth
	url  string
}

type response struct {
	Header     http.Header
	Body       io.ReadCloser
	StatusCode int
	Status     string
}

func checkHTTPStatus(resp response, expected ...int) error {
	for _, e := range expected {
		if resp.StatusCode == e {
			return nil
		}
	}
	return fmt.Errorf("http response status != %+v, got %d", expected, resp.StatusCode)
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
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return &response{}, fmt.Errorf("error creating POST request, %v", err)
	}

	if additionalHeaders == nil {
		additionalHeaders = make(map[string]string)
	}
	additionalHeaders[HEADER_X_TXC_PAYLOAD] = base64.StdEncoding.EncodeToString(bodyBytes)

	if c.auth != nil {
		h := hmac.New(sha256.New, []byte(c.auth.APISecret))
		h.Write(bodyBytes)
		signature := hex.EncodeToString(h.Sum(nil))
		additionalHeaders[HEADER_X_TXC_SIGNATURE] = signature
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
	thisHeaders["Content-type"] = "application/json"
	if c.auth != nil {
		thisHeaders[HEADER_X_TXC_APIKEY] = c.auth.APIKey
	}
	headers := mergeHeaders(additionalHeaders, thisHeaders)
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	resp, err := c.http.Do(request)
	if err != nil {
		fmt.Println(fmt.Sprintf("erro: %v", err))
		return nil, err
	}
	return &response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Header:     resp.Header,
		Body:       resp.Body,
	}, nil
}
