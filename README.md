# Mbase

## Development

### Start the application 


```bash
go run app.go
```

### Use local container

```
# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements

# Stop app in docker
make docker-stop

# Start app in docker
make docker-start
```

## Update swagger docs

1. Update comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
2. Generate docs files
```bash
swag init -g app.go
```

Go to http://localhost:3000\
Go to http://localhost:3000/swagger