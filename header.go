package yahc

import (
	"net/textproto"
)

// Headers to be passed along to the http request
type Header map[string][]string

// Adds a new header
func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

// Overwrites the given header, and sets it to the provided value
func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

// Removes the given key from the headers
func (h Header) Del(key string) {
	textproto.MIMEHeader(h).Del(key)
}

// Gets the given header
func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

// Creates a new header object based on the given map
func HeadersFromMap(m map[string]string) Header {
	header := Header{}
	for key, value := range m {
		header.Add(key, value)
	}
	return header
}
