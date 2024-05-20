// An example of a CRUD application.
//
// Run the app:
//
//	go run ./_examples/crud
//
// Create a post:
//
//	curl http://localhost:8080/posts --data-raw '{"data":{"title": "hello"}}'
//
// List all posts:
//
//	curl http://localhost:8080/posts
//
// Get a post:
//
//	curl http://localhost:8080/posts/0
//
// Update a post:
//
//	curl -X PATCH http://localhost:8080/posts/0 --data-raw '{"data":{"title": "world"}}'
//
// Delete a post:
//
//	curl -X DELETE http://localhost:8080/posts/0
package main

import (
	"fmt"
	"strconv"

	"github.com/orsinium-labs/josh"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var posts map[int]Post
var nextId = 0

func CreatePost(r josh.Req) josh.Resp[Post] {
	post, err := josh.Read[Post](r)
	if err != nil {
		respErr := josh.Error{Detail: err.Error()}
		return josh.BadRequest[Post](respErr)
	}
	post.ID = nextId
	posts[nextId] = post
	nextId += 1
	return josh.Created(post)
}

func ListPosts(r josh.Req) josh.Resp[[]Post] {
	list := make([]Post, 0)
	for _, post := range posts {
		list = append(list, post)
	}
	return josh.Ok(list)
}

func GetPost(r josh.Req) josh.Resp[Post] {
	postId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respErr := josh.Error{Detail: err.Error()}
		return josh.BadRequest[Post](respErr)
	}
	post, found := posts[postId]
	if !found {
		respErr := josh.Error{Detail: "post with the given ID does not exist"}
		return josh.NotFound[Post](respErr)
	}
	return josh.Ok(post)
}

func UpdatePost(r josh.Req) josh.Void {
	postId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respErr := josh.Error{Detail: err.Error()}
		return josh.BadRequest[josh.Z](respErr)
	}
	oldPost, found := posts[postId]
	if !found {
		respErr := josh.Error{Detail: "post with the given ID does not exist"}
		return josh.NotFound[josh.Z](respErr)
	}
	newPost, err := josh.Read[Post](r)
	if err != nil {
		respErr := josh.Error{Detail: err.Error()}
		return josh.BadRequest[josh.Z](respErr)
	}
	if newPost.Title != "" {
		oldPost.Title = newPost.Title
	}
	posts[postId] = oldPost
	return josh.NoContent[josh.Z]()
}

func DeletePost(r josh.Req) josh.Void {
	postId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respErr := josh.Error{Detail: err.Error()}
		return josh.BadRequest[josh.Z](respErr)
	}
	delete(posts, postId)
	return josh.NoContent[josh.Z]()
}

func main() {
	posts = make(map[int]Post)
	s := josh.NewServer(":8080")
	r := josh.Router{
		"/posts": {
			GET:  josh.Wrap(ListPosts),
			POST: josh.Wrap(CreatePost),
		},
		"/posts/{id}": {
			GET:    josh.Wrap(GetPost),
			PATCH:  josh.Wrap(UpdatePost),
			DELETE: josh.Wrap(DeletePost),
		},
	}
	r.Register(nil)
	fmt.Println("listening on http://localhost:8080")
	_ = s.ListenAndServe()
}
