package yahc

import (
	"testing"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func TestGet(t *testing.T) {
	out := make([]post, 0)
	err := Get("https://jsonplaceholder.typicode.com/posts", NoOptions, &out)
	if err != nil {
		t.Error(err)
	}

	if len(out) == 0 {
		t.Error("Got no content")
	}
}

func TestPost(t *testing.T) {
	newPost := post{UserID: 1, Title: "foo", Body: "The cake is a lie"}
	out := post{}
	err := Post("https://jsonplaceholder.typicode.com/posts", NoOptions, newPost, &out)
	if err != nil {
		t.Error(err)
	}

	if out.ID == 0 {
		t.Error("No post was created")
	}
}
