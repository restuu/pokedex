# POKEDEX

## Prerequisite

- [Docker](https://www.docker.com/products/docker-desktop/#)
- [docker-compose](https://docs.docker.com/compose/install/) (preinstalled with docker)
- [golang](https://go.dev/dl/)
- [golangci-lint](https://golangci-lint.run/usage/install/)

## Installation

```sh
docker-compose up pokedex
```

## Usage

Base URL: `http://localhost:3000`

### Enum

**Pokemon Type**

| Type     | Value |
|----------|-------|
| Normal   | 0     |
| Fighting | 1     |
| Flying   | 2     |
| Poison   | 3     |
| Ground   | 4     |
| Rock     | 5     |
| Bug      | 6     |
| Ghost    | 7     |
| Steel    | 8     |
| Fire     | 9     |
| Water    | 10    |
| Grass    | 11    |
| Electric | 12    |
| Psychic  | 13    |
| Ice      | 14    |
| Dragon   | 15    |
| Dark     | 16    |
| Fairy    | 17    |

### `POST /internal/pokemons`

Create new pokemon


**Request Body**
```json
{
    "name": "Pikachu",
    "type": 12,
    "image_url": ""
}
```

**Response Body**
```json
{
    "id": 1,
    "name": "Pikachu",
    "type": 12,
    "image_url": "",
    "created_at": "2023-02-20T15:13:08.579Z",
    "updated_at": "2023-02-20T15:13:08.579Z",
    "deleted_at": null
}
```

### `PUT /internal/pokemons/:id`

Update existing pokemon by id

**Request Body**
```json
{
    "id": 1,
    "name": "Pikachuuuu",
    "type": 10
}
```

**Response Body**
```json
{
    "id": 1,
    "name": "Pikachuuuu",
    "type": 10,
    "image_url": "",
    "created_at": "2023-02-20T15:13:08.579Z",
    "updated_at": "2023-02-20T18:16:17.522Z",
    "deleted_at": null
}
```

### `GET /pokemons`

Get all registered pokemons

**Response Body**
```json
[
    {
        "id": 1,
        "name": "Pikachu",
        "type": 10,
        "image_url": "",
        "created_at": "2023-02-20T15:13:08.579Z",
        "updated_at": "2023-02-20T15:13:08.579Z",
        "deleted_at": null
    }
]
```

## Notes

### Application

This application is build with `go` using `echo` webapi framework. All dependecies is being managed using `go.mod`

### Database of Choice

In this application database being used is MySQL with the consideration of complex joining and possibility of implementation of paging data which needs join operation

### Security

Admin should login first to get JWT token which will be needed in performing operation in `/internal/*` endpoint. However its not being implemented yet in this service due to time constraint

### Dependency Injection

Here Dependencies injection is being utilized using google wire library, however since its the first time for me using such library (previously I do it manually), the implementation might be a bit rough

### Design Pattern

Each component of this application such as user, pokemon, etc is grouped within it's domain (DDD). Each function layer has its own purpose such as
- _router_ for external communication
- _service_ for main business logic
- _repository_ for database connection layer

### Testing

Unfortunately this application is still lacking unit test coverage due to time constraint. However some reference on test method that I usually approach are
- making test file in the same package like `util.go, util_test.go` or make test package like `util.go,util_test/util_test.go`
- using testing library such as `testify` to make test suite with `vektra/mockery` to generate mock implementation