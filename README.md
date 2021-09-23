# buffalo-swagger

Buffalo middleware to automatically generate RESTful API documentation with Swagger 2.0.

[![Build Status](https://github.com/swaggo/buffalo-swagger/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/features/actions)
[![Codecov branch](https://img.shields.io/codecov/c/github/swaggo/buffalo-swagger/master.svg)](https://codecov.io/gh/swaggo/buffalo-swagger)
[![Go Report Card](https://goreportcard.com/badge/github.com/swaggo/buffalo-swagger)](https://goreportcard.com/report/github.com/swaggo/buffalo-swagger)
[![GoDoc](https://godoc.org/github.com/swaggo/buffalo-swagger?status.svg)](https://godoc.org/github.com/swaggo/buffalo-swagger)


## Usage

### Start using it
1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:
```sh
$ go get github.com/swaggo/swag/cmd/swag
```

3. The General API annotation lives in `actions/app.go`, run [Swag](https://github.com/swaggo/swag) in your Buffalo project root folder with the flag `-g actions/app.go`. [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```sh
$ swag init -g actions/app.go
```
4.Download [buffalo-swagger](https://github.com/swaggo/buffalo-swagger) by using:
```sh
$ go get -u github.com/swaggo/buffalo-swagger
```
And import following in your `actions/app.go` code, making sure to modify the last package name properly:

```go
import(
    buffaloSwagger "github.com/swaggo/buffalo-swagger"
    "github.com/swaggo/buffalo-swagger/swaggerFiles"
    _ "github.com/<github_name>/<project_name>/docs"
)
```

### Canonical example:
For a complete example take a look at the [example directory](https://github.com/swaggo/buffalo-swagger/tree/master/example)
Below you can find an extract from `actions/app.go`

```go
package actions

import(
    buffaloSwagger "github.com/swaggo/buffalo-swagger"
    "github.com/swaggo/buffalo-swagger/swaggerFiles"
    _ "github.com/swaggo/buffalo-swagger/example/docs"
)

[...]
var app *buffalo.App

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func App() *buffalo.App {
    if app == nil {
        app = buffalo.New(buffalo.Options{
            Env:          ENV,
            SessionStore: sessions.Null{},
            PreWares: []buffalo.PreWare{
                    cors.Default().Handler,
            },
            SessionName: "_example_session",
        })
    app.GET("/", HomeHandler)
    app.GET("/swagger/{doc:.*}", buffaloSwagger.WrapHandler(swaggerFiles.Handler))

}

return app
```

5. Run it, and browse to http://localhost:3000/swagger/index.html, you can see Swagger 2.0 Api documents.

![swagger_index.html](https://user-images.githubusercontent.com/8943871/36250587-40834072-1279-11e8-8bb7-02a2e2fdd7a7.png)

