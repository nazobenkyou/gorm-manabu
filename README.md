# GORM Manabu

For studying purpose and pocs

## Requirements

1. Docker compose
2. Docker
3. Golang 1.14

## Getting started

1. Execute the following commands

```bash
go mod download

go build

docker-compose up --build
```

2. Test the application

```bash
curl -XGET http://localhost:8080/tasks
```