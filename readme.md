# Go Vertical Slice Template

> A Golang boilerplate template, based on [Vertical Slice Architecture](https://jimmybogard.com/vertical-slice-architecture/) and [CQRS pattern](https://event-driven.io/en/cqrs_facts_and_myths_explained/) with using [Echo](https://github.com/labstack/echo), [Gorm](https://github.com/go-gorm/gorm), [Zap](https://github.com/go-gorm/gorm), [Viper](https://github.com/spf13/viper), [MediatR](https://github.com/mehdihadeli/Go-MediatR/) for CQRS and [uber-go/dig](https://github.com/uber-go/dig) for Dependency Injection.

**You can use this project as a template for building your backend application in Go, it's designed as a helpful starting point for your development.**

## Features

- ✅ Using `Vertical Slice Architecture` as a high level architecture
- ✅ Using `Data Centric Architecture` based on CRUD
- ✅ Using `CQRS Pattern` and `Mediator Pattern`on top of [mehdihadeli/Go-MediatR](https://github.com/mehdihadeli/Go-MediatR) library
- ✅ Using `Dependency Injection` and `Inversion of Control`on top of [uber-go/dig](https://github.com/uber-go/dig) library
- ✅ Using `RESTFul api` with [Echo](https://github.com/labstack/echo) framework and `Open-Api` using swagger with [swaggo/swag](https://github.com/swaggo/swag) library
- ✅ Using [go-playground/validator](https://github.com/go-playground/validator) for validating input data in the REST and gRpc
- ✅ Using `Gorm` and `SQLLite` for databases
- ✅ Using `Zap` for Logging
- ✅ Using `Viper` for configuration management

## Technologies - Libraries

- ✔️ **[`labstack/echo`](https://github.com/labstack/echo)** - High performance, minimalist Go web framework
- ✔️ **[`uber-go/zap`](https://github.com/uber-go/zap)** - Blazing fast, structured, leveled logging in Go.
- ✔️ **[`emperror/errors`](https://github.com/emperror/errors)** - Drop-in replacement for the standard library errors package and github.com/pkg/errors
- ✔️ **[`stretchr/testify`](https://github.com/stretchr/testify)** - A toolkit with common assertions and mocks that plays nicely with the standard library
- ✔️ **[`mehdihadeli/go-mediatr`](https://github.com/mehdihadeli/go-mediatr)** - Mediator pattern implementation in Golang and helpful in creating CQRS based applications.
- ✔️ **[`swaggo/swag`](https://github.com/swaggo/swag)** - Automatically generate RESTful API documentation with Swagger 2.0 for Go.
- ✔️ **[`go-gorm/gorm`](https://github.com/go-gorm/gorm)** - The fantastic ORM library for Golang, aims to be developer friendly
- ✔️ **[`go-playground/validator`](https://github.com/go-playground/validator)** - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving
- ✔️ **[`uber-go/dig`](https://github.com/uber-go/dig)** - A reflection based dependency injection toolkit for Go.
- ✔️ **[`spf13/viper`](https://github.com/spf13/viper)** - Go configuration with fangs
- ✔️ **[`caarlos0/env`](https://github.com/caarlos0/env)** - A simple and zero-dependencies library to parse environment variables into structs.
- ✔️ **[`joho/godotenv`](https://github.com/joho/godotenv)** - A Go port of Ruby's dotenv library (Loads environment variables from .env files)
- ✔️ **[`mcuadros/go-defaults`](https://github.com/mcuadros/go-defaults)** - Go structures with default values using tags

## Project Layout and Structure

projects structure is based on:

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## How to run the project?

We can run this Go boilerplate project with following steps:

- Clone this project.
- Move to your workspace: `cd your-workspace`.
- Clone this project into your workspace: `git clone https://github.com/mehdihadeli/go-vertical-slice-template`.
- Move to the project root directory: `cd go-vertical-slice-template`.
- Create a file `.env` similar to existing `.env` file at the root directory for your environment variables.
- Add application configurations based on enviroment (dev or production) in `config/config.development.json` or `config.production.json` files.
- [Install `go`](https://go.dev/doc/install) if not installed on your machine.
- Run `go run cmd/app/main.go`.
- Access API using [http://localhost:9080](http://localhost:9080).

## How to run the tests?

```bash
# Run all tests
go test ./...
```

## High Level Structure

```cmd
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   golangci.yml
│   readme.md
├───cmd
│   └───app
│           main.go
│
├───config
│       config.development.json
│       config.go
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
└───internal
    ├───catalogs
    │   ├───products
    │   │   │   mapper.go
    │   │   │
    │   │   ├───contracts
    │   │   │   │   endpoint.go
    │   │   │   │   product_respository.go
    │   │   │   │
    │   │   │   └───params
    │   │   │           product_route_params.go
    │   │   │
    │   │   ├───dtos
    │   │   │       product_dto.go
    │   │   │
    │   │   ├───features
    │   │   │   ├───creating_product
    │   │   │   │   ├───commands
    │   │   │   │   │       create_product.go
    │   │   │   │   │       create_product_handler.go
    │   │   │   │   │
    │   │   │   │   ├───dtos
    │   │   │   │   │       create_product_request_dto.go
    │   │   │   │   │       create_product_response.go
    │   │   │   │   │
    │   │   │   │   ├───endpoints
    │   │   │   │   │       create_product_endpoint.go
    │   │   │   │   │
    │   │   │   │   └───events
    │   │   │   │           product_created.go
    │   │   │   │           product_created_handler.go
    │   │   │   │
    │   │   │   └───getting_product_by_id
    │   │   │       ├───dtos
    │   │   │       │       get_product_by_id_request_dto.go
    │   │   │       │       get_product_by_id_response.go
    │   │   │       │
    │   │   │       ├───endpoints
    │   │   │       │       get_product_by_id_endpoint.go
    │   │   │       │
    │   │   │       └───queries
    │   │   │               get_product_by_id.go
    │   │   │               get_product_by_id_handler.go
    │   │   │
    │   │   ├───models
    │   │   │       product.go
    │   │   │
    │   │   └───repository
    │   │           inmemory_product_repository.go
    │   │
    │   └───shared
    │       ├───app
    │       │   ├───application
    │       │   │       application.go
    │       │   │       application_endpoints.go
    │       │   │       application_mediatr.go
    │       │   │       application_migration.go
    │       │   │
    │       │   └───application_builder
    │       │           application_builder.go
    │       │           application_builder_dependencies.go
    │       │
    │       └───behaviours
    │               request_logger_behaviour.go
    │
    └───pkg
        ├───config
        │   │   config_helper.go
        │   │   dependency.go
        │   │
        │   └───environemnt
        │           environment.go
        │
        ├───constants
        │       constants.go
        │
        ├───database
        │   │   db.go
        │   │   dependency.go
        │   │
        │   └───options
        │           gorm_options.go
        │
        └───reflection
            └───type_mappper
                    type_mapper.go
                    type_mapper_test.go
                    unsafe_types.go

```

## Application Structure

In this project I used [vertical slice architecture](https://jimmybogard.com/vertical-slice-architecture/) or [Restructuring to a Vertical Slice Architecture](https://codeopinion.com/restructuring-to-a-vertical-slice-architecture/) also I used [feature folder structure](http://www.kamilgrzybek.com/design/feature-folders/) in this project.

- We treat each request as a distinct use case or slice, encapsulating and grouping all concerns from front-end to back.
- When We are adding or changing a feature in an application in n-tire architecture, we are typically touching many different "layers" in an application. we are changing the user interface, adding fields to models, modifying validation, and so on. Instead of coupling across a layer, we couple vertically along a slice and each change affects only one slice.
- We `Minimize coupling` `between slices`, and `maximize coupling` `in a slice`.
- With this approach, each of our vertical slices can decide for itself how to best fulfill the request. New features only add code, we're not changing shared code and worrying about side effects. For implementing vertical slice architecture using cqrs pattern is a good match.

![](./assets/vertical-slice-architecture.jpg)

![](./assets/vsa2.png)

Also here I used [CQRS](https://www.eventecommerce.com/cqrs-pattern) for decompose my features to very small parts that makes our application:

- maximize performance, scalability and simplicity.
- adding new feature to this mechanism is very easy without any breaking change in other part of our codes. New features only add code, we're not changing shared code and worrying about side effects.
- easy to maintain and any changes only affect on one command or query (or a slice) and avoid any breaking changes on other parts
- it gives us better separation of concerns and cross-cutting concern (with help of MediatR behavior pipelines) in our code instead of a big service class for doing a lot of things.

With using [CQRS](https://event-driven.io/en/cqrs_facts_and_myths_explained/), our code will be more aligned with [SOLID principles](https://en.wikipedia.org/wiki/SOLID), especially with:

- [Single Responsibility](https://en.wikipedia.org/wiki/Single-responsibility_principle) rule - because logic responsible for a given operation is enclosed in its own type.
- [Open-Closed](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle) rule - because to add new operation you don’t need to edit any of the existing types, instead you need to add a new file with a new type representing that operation.

Here instead of some [Technical Splitting](http://www.kamilgrzybek.com/design/feature-folders/) for example a folder or layer for our `services`, `controllers` and `data models` which increase dependencies between our technical splitting and also jump between layers or folders, We cut each business functionality into some vertical slices, and inner each of these slices we have [Technical Folders Structure](http://www.kamilgrzybek.com/design/feature-folders/) specific to that feature (command, handlers, infrastructure, repository, controllers, data models, ...).

Usually, when we work on a given functionality we need some technical things for example:

- API endpoint (Controller)
- Request Input (Dto)
- Request Output (Dto)
- Some class to handle Request, For example Command and Command Handler or Query and Query Handler
- Data Model

Now we could all of these things beside each other, and it decreases jumping and dependencies between some layers or folders.

Keeping such a split works great with CQRS. It segregates our operations and slices the application code vertically instead of horizontally. In Our CQRS pattern each command/query handler is a separate slice. This is where you can reduce coupling between layers. Each handler can be a separated code unit, even copy/pasted. Thanks to that, we can tune down the specific method to not follow general conventions (e.g. use custom SQL query or even different storage). In a traditional layered architecture, when we change the core generic mechanism in one layer, it can impact all methods.

## Live Reloading In Development

For live reloading in dev mode I use [air](https://github.com/cosmtrek/air) library. for guid about using this tools you can [read this article](https://mainawycliffe.dev/blog/live-reloading-golang-using-air/).

For running app in `live reload mode`, inner type bellow command after [installing air](https://github.com/cosmtrek/air?ref=content.mainawycliffe.dev#via-go-install):

```bash
air
```
