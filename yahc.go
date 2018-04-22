package yahc

import (
	"net/http"
)

// Gets a new client, based on the given client
// use this if you want provide e.g. a cookie jar
// for interacting with stateful services
func NewClient(inner *http.Client) *Client {
	return &Client{client: inner}
}

var (
	// A precreated client. Uses the http.Default client
	// All direct yahc.Get/Post/Put/Delete/Patch methods uses this
	// client.
	//
	// If another client should be used instead, request a custom
	// client instead using yahc.NewClient
	DefaultClient = NewClient(http.DefaultClient)

	// A default set of options
	// Will treat everything as json
	NoOptions = RequestOptions{
		Headers: HeadersFromMap(map[string]string{ContentType: ApplicationJsonContentType}),
	}
)

// Does a GET request using the default client
func Get(url string, options RequestOptions, v interface{}) error {
	return DefaultClient.Get(url, options, v)
}

// Does a POST request using the default client
func Post(url string, options RequestOptions, body, v interface{}) error {
	return DefaultClient.Post(url, options, body, v)
}

// Does a PUT request using the default client
func Put(url string, options RequestOptions, body, v interface{}) error {
	return DefaultClient.Put(url, options, body, v)
}

// Does a delete request using the default client
func Delete(url string, options RequestOptions, v interface{}) error {
	return DefaultClient.Delete(url, options, v)
}
