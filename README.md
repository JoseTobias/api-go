# Simple go api

This project is a simple api to learning the first steps on golang.

# Environments

To set the environments variables you must to create a `.env` file. The only variable that can be defined is PORT, that is the port that api listen.

# Run

To run project you need to install godotenv so run the follows commands to run the project.

```
go mod init api-go
go run main.go
```

So do GET to `localhost:3800` or if you selected PORT in `.env` do GET to `localhost:PORT` to test the api.