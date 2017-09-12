# go-cms-api

This is a small api service provides getting, creating posts via REST. It also includes the JWT authentication to prevent anonymous put the data into database.

## How to use

### Get the project 

Assume that you've already configured your GOPATH.

```
go get github.com/tatthien/go-cms-api
```


### Configure app information
Copy or rename the file `.env.example` to `.env`. Then enter your own server ip, application port, database information.

```
APP_IP=your_server_ip
APP_PORT=your_application_port
DATABASE_NAME=your_database_name
DATABASE_USERNAME=your_datatase_username
DATABASE_PASSWORD=your_database_password
```

If you folk this project, make sure you do not push this `.env` file to Github. I think you know why ;)

### Configure token

This project is using `jwt-golang` package to generate and validate the token.

## Install database

```
go run install.go
```

## Start server

```
go run main.go
```

## Endpoints

### Authentication

Endpoint: `POST /api/v1/login`

Request data:

- `email`: Admin email
- `password`: Admin password

### Get list of posts

Endpoint: `GET /api/v1/posts`

Request data:

- `page`: default is `1`. Current page.
- `limit`: default is `10`. Litmit the posts from response.
- `post_type`: (`post`|`til`) default is `post`. Get the specific post type.
