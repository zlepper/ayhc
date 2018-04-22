# Yet another HTTP client (yahc)

yach is a very simple opinionated http client, that takes most of the decisions 
out of the hands of the developer, and tries to figure out what is actually 
supposed to be going on. 

yach comes with build in support for interfacing with REST endpoints, but can 
very easily be extended. In fact the standard json support is an extension 
build into the library. 

## Installation

```
go get github.com/zlepper/yahc
```

Godoc: https://godoc.org/github.com/zlepper/ayhc

## Examples
Here are some basic examples of how this library can be used. 

### Getting a resource
```go
package main

import "github.com/zlepper/yahc"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	out := make([]Post, 0)
	err := yahc.Get("https://jsonplaceholder.typicode.com/posts", yahc.NoOptions, &out)
	// Response is now in 'out' if the request didn't fail
}
```

### Posting a resource
```go
package main

import "github.com/zlepper/yahc"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	newPost := Post{UserID:1,Title:"foo",Body:"The cake is a lie"}
	out := Post{}
	err := yahc.Post("https://jsonplaceholder.typicode.com/posts", yahc.NoOptions, newPost, &out)
	// Response is now in 'out' if the request didn't fail
}
```

Same pattern follows for PUT and DELETE

## Advanced usage
Advanced usage is for when you are not just working with json, or might require a custom client

### Custom client
A new client can be created by doing `yahc.NewClient(&http.Client{})`. 
Example use case for this adding a cookie-jar to the client to work
with stateful apis

```go
package main

import (
    "net/http"
    "net/http/cookiejar"
    "github.com/zlepper/yahc"
)

func main() {
    cookieJar, _ := cookiejar.New(nil)
    
    httpClient := &http.Client{
        Jar: cookieJar,
    }
    
    client := yahc.NewClient(httpClient)
}
```

### Custom parses
yahc attempts to automatically parse your request and response into 
something that can be send to servers. All these parses can be changed, 
and additional ones registered to allow for working with more types of
content. By default only a JSON parser is provided.

To register a custom parser, implement the 
```go
// A request parser will take to the request body struct, and turn it into
// a byte array that can be passed along to the actual request
type RequestParser interface {
	// Should convert the given argument into a byte array
	Convert(v interface{}) ([]byte, error)
}
```

or 
```go
// A ResponseParser should convert a response body into an object
type ResponseParser interface {
	// Should convert the r into v
	// Closing the reader is handled by yahc
	Convert(r io.Reader, v interface{}) error
}
```

depending on your use-case, and register them with 
`RegisterRequestParser` or `RegisterResponseParser`. 
A parser can be registered for multiple content types. 

For an example, check the [json-parser.go](json-parser.go) file,
which is what handles the build in json parser.
