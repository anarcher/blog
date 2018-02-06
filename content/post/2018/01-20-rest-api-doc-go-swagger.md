+++
title="Make REST API Documentation using swagger in Go"
tags=["tool","documentation"]
draft=false
date="2018-01-20"
+++

For golang based HTTP/REST API documentation,I choose [swagger][]. [go-swagger][] has several features for swagger documentation. The go-swagger can generate swagger spec based code generation but I already have an REST API server. I use go-swagger with golang comment annotation for swagger spec generation.

For REST API development, Design first with writing spec and then generating codes from it is a good approach. [goa][] is a famous tool for this style.  

# Generate go-swagger spec

The go-swagger has a command that will let you generate a swagger spec document from codes. The command scan all files from current directory or specific file path to sub directories.   

To generate a spec:

`swagger generate spec` command makes swagger.json spec. 

```
// Package rest HTTP API.
//
// Schemes: http
// Host: localhost
// BasePath: /
// Version: 0.0.1
// Consumes:
// - application/json
// Produces:
// - application/json
// swagger:meta
package rest
```

```
//go:generate swagger generate spec -m -i ./swagger.yml -o ./static/swagger.json
```

# Operations and struct model mapping

One of reasons why I choose swagger or api documentation generation tool is that parsing source code and using on api spec. 

```
// swagger:operation POST /v1/articles Article create_article
// Create an article resource
//
// ---
// summary: Create article
// parameters:
//  - in: body
//    required: true
//    schema:
//       $ref: "#/definitions/createArticleRequest"
// responses:
// responses:
//  '200':
//   description: successful operation
//   schema:
//    $ref: "#/definitions/Article"
//  '500':
//   description: failed operation (error)
//   schema:
//    $ref: "#/definitions/Error"
//
func createAppHandler(context *gin.Context) {
...
```

The `definitions/Article` and `definitions/Error` are mapping from struct with `swagger:model`

```
// swagger:model 
type Article struct {
    ID         string    `db:"id" json:"id" yaml:"id"`
    Name       string    `db:"name" json:"name" yaml:"name"`
    CreateTime time.Time `db:"create_time" json:"create_time"`
    UpdateTime time.Time `db:"update_time" json:"update_time"`
}
```

It is convenient if you have complex request and response payloads.


# Embedded static asset(swagger.json) with go-bindata

I want to embedded the swagger-ui and [redoc][] frontend page to my rest api server. [go-bindata][] converts any text or binary file into Go source code.

```
go:generate go-bindata -prefix "static/" -o static/bindata.go -pkg static static/...
```

That's it. the `swagger.json` is embedded in golang code. With gin framework, you can use it with [staticbin][].

```
import  "github.com/olebedev/staticbin"

...

router.Use(staticbin.Static(static.Asset, staticbin.Options{
        Dir: "/static",
}))

```

Embedded with swagger-ui, you can use [swagger-doc][]'s assets for it. 

```
import (
    swaggerDoc "github.com/utahta/swagger-doc"
    swaggerAsset "github.com/utahta/swagger-doc/assets"
)


...

router.Use(staticbin.Static(swaggerAsset.Asset, staticbin.Options{
        Dir: "/swagger-ui",
}))
router.Any("/redoc", gin.WrapH(swaggerDoc.NewRedocHandler("/static/swagger.json", "redoc")))
```


I'm not sure this approach is good. I think that better approach exists for it. [apiblueprint][] is simple yaml structure than swagger's yaml. but I have not found good [apiblueprint][] utils for my taste for golang code generation or spec generation from code.


[swagger]: https://swagger.io/
[go-swagger]: https://goswagger.io/
[redoc]: https://github.com/Rebilly/ReDoc
[goa]: https://goa.design/
[go-bindata]: https://github.com/jteeuwen/go-bindata
[staticbin]: https://github.com/olebedev/staticbin
[swagger-doc]: https://github.com/utahta/swagger-doc
[apiblueprint]: https://apiblueprint.org/

