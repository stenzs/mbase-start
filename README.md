# Mbase

## Development

### Start the application 


```bash
make docker-start
```

### Use local container

```
# Clean upload folder
make clean-uploads

# Start app in docker
make docker-start

# Stop app in docker
make docker-stop

# Remove app docker containers
make docker-rm
```

## Update swagger docs

1. Update comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
2. Generate docs files
```bash
swag init -g app.go
```

Go to http://localhost:3000\
Go to http://localhost:3000/swagger