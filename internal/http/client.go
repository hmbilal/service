package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/liamylian/jsontime"
)

var json = jsontime.ConfigWithCustomTimeFormat

type Client interface {
	DoRequest(req *Request) (*http.Response, error)
}

type client struct {
	baseURI    string
	baseClient *http.Client
}

type BasicAuth struct {
	Username, Password string
}

type Request struct {
	Method      string
	Path        string
	Body        []byte
	QueryParams map[string]string
	Headers     map[string]string
	BasicAuth   *BasicAuth
}

func NewClient(baseURI string) Client {
	return &client{
		baseURI: baseURI,
		baseClient: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
	}
}

func (c *client) DoRequest(req *Request) (*http.Response, error) {
	httpRequest, err := http.NewRequestWithContext(context.Background(), req.Method, c.buildURL(req.Path), bytes.NewBuffer(req.Body))
	if err != nil {
		return nil, err
	}

	buildHeaders(httpRequest, req.Headers)
	buildAuth(httpRequest, req.BasicAuth)
	buildQueryParams(httpRequest, req.QueryParams)

	response, err := c.baseClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if !isResponseStatusSuccess(response) {
		return response, &RequestError{StatusCode: response.StatusCode, Body: response.Body}
	}

	return response, nil
}

func DecodeResponse(response *http.Response, value interface{}) error {
	defer func() {
		_ = response.Body.Close()
	}()

	return json.NewDecoder(response.Body).Decode(value)
}

func (c *client) buildURL(path string) string {
	return fmt.Sprintf("%s/%s", c.baseURI, path)
}

func buildHeaders(httpRequest *http.Request, headers map[string]string) {
	for key, value := range headers {
		httpRequest.Header.Set(key, value)
	}
}

func buildAuth(httpRequest *http.Request, auth *BasicAuth) {
	if auth != nil && auth.Username != "" && auth.Password != "" {
		httpRequest.SetBasicAuth(auth.Username, auth.Password)
	}
}

func buildQueryParams(httpRequest *http.Request, params map[string]string) {
	query := httpRequest.URL.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	httpRequest.URL.RawQuery = query.Encode()
}

func isResponseStatusSuccess(response *http.Response) bool {
	return response.StatusCode/100 == 2
}
