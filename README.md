# Go REST Demo

Creating a web server is one of the fields where Go truly shines. This repo demonstrates three ways you can create REST APIs with Go.

All the details can be found in the related articles:
* [Go REST Guide. The Standard Library](https://www.jetbrains.com/go/guide/tutorials/rest_api_series/stdlib/)
* [Go REST Guide. gorilla/mux Router](https://www.jetbrains.com/go/guide/tutorials/rest_api_series/gorilla-mux/)
* [Go REST Guide. Gin Framework](https://www.jetbrains.com/go/guide/tutorials/rest_api_series/gin/)

## Demo Application

The demo application is a simple REST API for managing a collection of recipes. It showcases how to implement basic CRUD (Create, Read, Update, Delete) operations using different Go frameworks.

The application uses an in-memory data store, so the recipes will be lost when the application is stopped.

## API Endpoints

The API provides the following endpoints:

| Method   | Path             | Description                  |
|----------|------------------|------------------------------|
| `GET`    | `/recipes`       | Get a list of all recipes    |
| `POST`   | `/recipes`       | Create a new recipe          |
| `GET`    | `/recipes/{id}`  | Get a single recipe by its ID|
| `PUT`    | `/recipes/{id}`  | Update an existing recipe    |
| `DELETE` | `/recipes/{id}`  | Delete a recipe by its ID    |

This structure can be used as a template for adding new frameworks to this showcase.

## How to Run

To run any of the versions, navigate to the corresponding directory in the `cmd` folder and use `go run`:

### Standard Library
```bash
go run cmd/standardlib/main.go
```

### gorilla/mux
```bash
go run cmd/gorilla/main.go
```

### Gin
```bash
go run cmd/gin/main.go
```
