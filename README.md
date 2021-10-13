# Go GraphQL

Simple vanilla [Go](https://golang.org/) CRUD application with [mongoDB](https://www.mongodb.com/) database with its [mflix](https://github.com/neelabalan/mongodb-sample-dataset/tree/main/sample_mflix) dataset that I use for my thesis about benchmarking [REST API](https://restfulapi.net/) and [GraphQL](https://graphql.org/).

## Usage

To use this application run

```bash
go run main.go
```

The server will be served at `http://localhost:8081`

This app provides the following endpoints:

* `GET /movies`: return all movies data with limit from query string (default: 10)
* `POST /comments`: creates new comment
* `PUT /comments/:id`: updates an existing comment
* `DELETE /comments/:id`: deletes a comment

## Related Repository

Below is another repository used for my thesis.

* [Go GraphQL](https://github.com/adrianedy/go-graphql)
* [PHP REST](https://github.com/adrianedy/php-rest)
* [PHP GraphQL](https://github.com/adrianedy/php-graphql)