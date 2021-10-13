package router

import (
	"github.com/adrianedy/go-rest/comments"
	"github.com/adrianedy/go-rest/movies"
)

var routes = []route{
	post("/comments", comments.Create),
	put("/comments/:id", comments.Update),
	delete("/comments/:id", comments.Delete),
	get("/movies", movies.Read),
}
