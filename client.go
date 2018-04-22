package yahc

import (
	"bytes"
	"io"
	"net/http"
	nurl "net/url"
)

// A client that can interact with external services
type Client struct {
	client *http.Client
}

// Headers and query parameters for the request
type RequestOptions struct {
	// Headers to attach to the request
	Headers     Header
	// Query parameters to add to the request
	QueryParams QueryParam
}

// Internal utility method for moving the query parameters to a url
func addQueryParamsToUrl(url *nurl.URL, params QueryParam) {
	// Pass along the query parameters
	for key, values := range params {
		for _, value := range values {
			url.Query().Add(key, value)
		}
	}
}

// Internal utility method for adding the headers to the request
func addHeadersToRequest(request *http.Request, header Header) {
	for key, values := range header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
}

// Attempts to detect the content type of the response
// As this function consumes the first 512 bytes of the response
// a new reader is returned, that also contains the start bytes
func detectContentType(r io.Reader) (string, io.Reader, error) {

	data := make([]byte, 512)
	read, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", nil, err
	}

	data = data[0:read]

	contentType := http.DetectContentType(data)

	// Get a reader that contains the initial content too
	newReader := io.MultiReader(bytes.NewReader(data), r)

	return contentType, newReader, nil
}

// Automatically detect the content type in the response
// if it cannot be determined, returns "application/octet-stream"
// The reader returned from this function should be used instead
// of the response.Body, as some data might have been consumed
// during detection
func detectContentTypeFromResponse(response *http.Response) (contentType string, reader io.Reader, err error) {
	reader = response.Body

	contentType = response.Header.Get(ContentType)
	if contentType == "" {
		contentType, reader, err = detectContentType(reader)
	}

	return contentType, reader, err
}

// Automatically parses the response
func autoParseResponse(response *http.Response, v interface{}) error {
	contentType, reader, err := detectContentTypeFromResponse(response)
	if err != nil {
		return err
	}
	return ParseResponse(contentType, reader, v)
}

func (c *Client) doWithoutBody(method, url string, options RequestOptions, v interface{}) error {
	u, err := nurl.Parse(url)
	if err != nil {
		return err
	}

	addQueryParamsToUrl(u, options.QueryParams)

	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	addHeadersToRequest(request, options.Headers)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if v != nil {
		err = autoParseResponse(response, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) doWithBody(method, url string, options RequestOptions, body interface{}, v interface{}) error {
	u, err := nurl.Parse(url)
	if err != nil {
		return err
	}

	addQueryParamsToUrl(u, options.QueryParams)

	contentType := options.Headers.Get(ContentType)

	data, err := ConvertRequestBody(contentType, body)
	if err != nil {
		return err
	}

	options.Headers.Set(ContentType, contentType)

	request, err := http.NewRequest(method, u.String(), bytes.NewReader(data))
	if err != nil {
		return err
	}

	addHeadersToRequest(request, options.Headers)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if v != nil {
		err = autoParseResponse(response, v)
		if err != nil {
			return err
		}
	}

	return nil
}

// Does a get request to the given resource
// v should be a pointer to a struct that should be filled with values
func (c *Client) Get(url string, options RequestOptions, v interface{}) error {
	return c.doWithoutBody(http.MethodGet, url, options, v)
}

// Does a post request to the given resource
// Body of the struct that should be send to the resource
// v if a pointer to an output, in case a response is expected
func (c *Client) Post(url string, options RequestOptions, body interface{}, v interface{}) error {
	return c.doWithBody(http.MethodPost, url, options, body, v)
}

// Does a put request to the given resource
func (c *Client) Put(url string, options RequestOptions, body interface{}, v interface{}) error {
	return c.doWithBody(http.MethodPut, url, options, body, v)
}

// Does a delete to the given resource
func (c *Client) Delete(url string, options RequestOptions, v interface{}) error {
	return c.doWithoutBody(http.MethodDelete, url, options, v)
}
