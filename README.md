# goCleanArchitecture

### 1. Gerar arquivo proto:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/order.proto
```

### 2. Go mod tidy
```bash
go get github.com/99designs/gqlgen/codegen/config@v0.17.55
go get github.com/99designs/gqlgen/internal/imports@v0.17.55
go get github.com/99designs/gqlgen/internal/imports@v0.17.55
go get github.com/99designs/gqlgen@v0.17.55
go mod tidy
```

### 3. Gerar gqlgen
```bash
go run github.com/99designs/gqlgen generate
```

### 4. Inicie o Docker

```bash
docker-compose down
docker-compose up --build
```
