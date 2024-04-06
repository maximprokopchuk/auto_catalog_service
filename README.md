# Auto Reference Microservice

## Prerequisites

- Protobuf - https://protobuf.dev/
- GRPC - https://grpc.io/docs/languages/go/quickstart/
- SQLC - https://sqlc.dev/

## Usage

Create `configs/config.toml` file. You can use `configs/config.example.toml` as example

Install dependencies:

``` bash:
make deps
```

Run server:

``` bash:
make run

```

Build binary:

``` bash:
make build
```

## Development

Update SQLC schema (`schema.sql`) and queries (`query.sql`). Then generate Go code (models, interfaces etc):

``` bash:
 make sqlc_generate
```

Update protobuf spec (`proto/auto_reference_catalog.proto`) and then generate Go code:

``` bash:
 make proto_generate
```

GRPC procedures implementation is in `internal/grpcserver/grpcserver.go`

## Tests

Run tests:

``` bash:
make test
```

## Linter

Run linter:

``` bash:
make lint
```
