# go-cms-api

## Get the project

Assume that you've already configured your GOPATH.

```
go get github.com/tatthien/go-cms-api
```

## Configuration

### 1. Configure database
Copy or rename the file `.env.example` to `.env`. Then enter your own server ip, application port, database information.

```
APP_IP = your server ip
APP_PORT = your application port
DATABASE_NAME = your database name
DATABASE_USERNAME = your datatase username
DATABASE_PASSWORD = your database password
```

If you folk this project, make sure you do not push this `.env` file to Github. I think you know why ;)

### 2. Configure token

This project is using `jwt-golang` package to generate and validate the token.

## Install database

```
go run install.go
```

## Start server

```
go run main.go
```

## Run with your custom domain
